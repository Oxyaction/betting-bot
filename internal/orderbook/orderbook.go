package orderbook

import (
	"fmt"
	"sync"

	"gitlab.com/fireferretsbet/tg-bot/internal/user"
	"github.com/google/uuid"
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

type OrderBook struct {
	mu sync.RWMutex
	ob *ob.OrderBook

	orderMap   OrderMap
	user2order map[*user.User][]uuid.UUID
	order2user map[uuid.UUID]*user.User
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

func (b *OrderBook) Cancel(user *user.User, orderID uuid.UUID) (err error) {
	b.mu.Lock()
	b.ob.CancelOrder(orderID)

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
		order := b.orderMap[o.ID()]
		order.Status = OrderStatusMatched
		order.Matched = order.Qty
		order.Unmatched = decimal.Zero
		b.orderMap[o.ID()] = order
		changed = append(changed, order)
	}

	if partial != nil {
		o := b.orderMap[partial.ID()]
		o.Unmatched = partial.Quantity()
		o.Matched = o.Qty.Sub(o.Unmatched)
		order.Status = OrderStatusPartial
		b.orderMap[partial.ID()] = o
		changed = append(changed, o)
	}

	return
}

// TODO:
func (b *OrderBook) notifyUsers(changed []Order) {
	b.mu.RLock()
	b.mu.RUnlock()
}
