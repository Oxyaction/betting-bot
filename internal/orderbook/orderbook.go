package orderbook

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"

	"github.com/google/uuid"
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type OrderBook struct {
	log *logrus.Logger
	mu  sync.RWMutex
	ob  *ob.OrderBook

	orderMap   OrderMap
	user2order map[*user.User][]uuid.UUID
	order2user map[uuid.UUID]*user.User
}

func NewOrderBook(log *logrus.Logger) *OrderBook {
	return &OrderBook{
		log:        log,
		ob:         ob.NewOrderBook(),
		orderMap:   OrderMap{},
		user2order: map[*user.User][]uuid.UUID{},
		order2user: map[uuid.UUID]*user.User{},
	}
}

func (b *OrderBook) Place(user *user.User, order Order) (changedMap map[*user.User][]Order, err error) {
	b.mu.Lock()
	b.orderMap[order.ID] = order
	b.order2user[order.ID] = user
	b.user2order[user] = append(b.user2order[user], order.ID)
	changed, err := b.place(user, order)
	b.mu.Unlock()
	if err != nil {
		return
	}
	changedMap = b.makeChangedMap(changed)

	return
}

func (b *OrderBook) Cancel(user *user.User, orderID uuid.UUID) (order Order, err error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	memUser, order2userOk := b.order2user[orderID]
	order, orderMapOk := b.orderMap[orderID]
	switch {
	case !order2userOk || memUser.ID() != user.ID():
		err = ErrOrderInvalidOrderUser
	case !orderMapOk:
		err = ErrOrderNotFound
	case order.Status == OrderStatusMatched:
		err = ErrOrderAlreadyMatched
	case order.Status == OrderStatusCanceled:
		err = ErrOrderAlreadyCanceled
	}
	if err != nil {
		return
	}

	o := b.ob.CancelOrder(orderID)
	order.Status = OrderStatusCanceled
	order.Unmatched = o.Quantity()
	order.Matched = order.Qty.Sub(order.Unmatched)
	b.orderMap[orderID] = order

	b.mu.Unlock()
	return
}

func (b *OrderBook) place(
	user *user.User,
	order Order,
) (
	changed []Order,
	err error,
) {
	setPlacedStatus := true
	// TODO: tx
	balance := user.GetBalance()
	if balance.LessThan(order.Qty) {
		err = ErrNotEnoughFounds
		return
	}

	done, partial, _, err := b.ob.ProcessLimitOrder(order.Side, order.ID, order.Qty, order.Coeff)
	if err != nil {
		return
	}

	_, err = user.ChangeBalance(
		fmt.Sprintf("PlaceOrder %s: coeff: %s, qty %s", order.Side, order.Coeff, order.Qty),
		order.Qty.Neg(),
	)

	for _, o := range done {
		memOrder := b.orderMap[o.ID()]
		memOrder.Status = OrderStatusMatched
		memOrder.Matched = memOrder.Qty
		memOrder.Unmatched = decimal.Zero
		b.orderMap[o.ID()] = memOrder
		changed = append(changed, memOrder)
		if memOrder.ID == order.ID {
			setPlacedStatus = false
		}
	}

	if partial != nil {
		memOrder := b.orderMap[partial.ID()]
		memOrder.Unmatched = partial.Quantity()
		memOrder.Matched = memOrder.Qty.Sub(memOrder.Unmatched)
		memOrder.Status = OrderStatusPartial
		b.orderMap[partial.ID()] = memOrder
		changed = append(changed, memOrder)
		if memOrder.ID == order.ID {
			setPlacedStatus = false
		}
	}

	if setPlacedStatus {
		order.Status = OrderStatusPlaced
		b.orderMap[order.ID] = order
		changed = append(changed, order)
	}

	return
}

func (b *OrderBook) makeChangedMap(changed []Order) (changedMap map[*user.User][]Order) {
	changedMap = map[*user.User][]Order{}
	b.mu.RLock()
	for _, order := range changed {
		u := b.order2user[order.ID]
		changedMap[u] = append(changedMap[u], order)
	}
	b.mu.RUnlock()
	return
}

func (b *OrderBook) Unmatched() (lay, back []*ob.PriceLevel) {
	b.mu.RLock()
	lay, back = b.ob.Depth()
	b.mu.RUnlock()

	return
}

func (b *OrderBook) Matched() (lay, back []*ob.PriceLevel) {
	mBack := map[string]ob.PriceLevel{}
	mLay := map[string]ob.PriceLevel{}
	b.mu.RLock()
	for _, o := range b.orderMap {
		coeff := o.Coeff.String()
		if o.Side == ob.Back {
			z := mBack[coeff]
			z.Price = o.Coeff
			z.Quantity = z.Quantity.Add(o.Matched)
			mBack[coeff] = z
		} else {
			z := mLay[coeff]
			z.Price = o.Coeff
			z.Quantity = z.Quantity.Add(o.Matched)
			mLay[coeff] = z
		}
	}
	b.mu.RUnlock()

	for _, pl := range mBack {
		back = append(back, &pl)
	}

	for _, pl := range mLay {
		lay = append(lay, &pl)
	}

	return
}
