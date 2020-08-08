package orderbook

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type OrderBook struct {
	mu sync.RWMutex
	ob *ob.OrderBook

	orderMap   OrderMap
	user2order map[*user.User][]uuid.UUID
	order2user map[uuid.UUID]*user.User
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		ob:         ob.NewOrderBook(),
		orderMap:   OrderMap{},
		user2order: map[*user.User][]uuid.UUID{},
		order2user: map[uuid.UUID]*user.User{},
	}
}

func (b *OrderBook) Place(user *user.User, order Order) (err error) {
	b.mu.Lock()
	b.orderMap[order.ID] = order
	b.order2user[order.ID] = user
	b.user2order[user] = append(b.user2order[user], order.ID)
	changed, err := b.place(user, order)
	b.mu.Unlock()
	if err != nil {
		return
	}

	// TODO
	b.notifyUsers(changed)

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
		fmt.Sprintf("Place %s: coeff: %s, qty %s", order.Side, order.Coeff, order.Qty),
		order.Qty.Neg(),
	)

	for _, o := range done {
		memOrder := b.orderMap[o.ID()]
		memOrder.Status = OrderStatusMatched
		memOrder.Matched = order.Qty
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

// TODO:
func (b *OrderBook) notifyUsers(changed []Order) {
	b.mu.RLock()
	b.mu.RUnlock()
}
