package orderbook

import (
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func (b *OrderBook) SettleDraw(reason string) {
	b.mu.Lock()
	b.settleDraw(reason)
	b.mu.Unlock()
}

func (b *OrderBook) SettleBack(reason string) {
	b.mu.Lock()
	b.settle(reason, ob.Back)
	b.mu.Unlock()
}

func (b *OrderBook) SettleLay(reason string) {
	b.mu.Lock()
	b.settle(reason, ob.Lay)
	b.mu.Unlock()
}

func (b *OrderBook) settle(reason string, side ob.Side) {
	for user, orderIDs := range b.user2order {
		var totalDiff decimal.Decimal
		for _, orderID := range orderIDs {
			order := b.orderMap[orderID]
			if order.Matched.IsZero() {
				continue
			}
			if order.Side == side {
				totalDiff = totalDiff.Add(order.Settle(side))
			}
		}
		reason += ", " + side.String() + ": "
		_, err := user.ChangeBalance(reason, totalDiff)
		if err != nil {
			log.Error("user.ChangeBalance", reason, err)
		}
	}
}

func (b *OrderBook) settleDraw(reason string) {
	for user, orderIDs := range b.user2order {
		var totalDiff decimal.Decimal
		for _, orderID := range orderIDs {
			order := b.orderMap[orderID]
			if order.Status == OrderStatusCanceled {
				totalDiff = totalDiff.Add(order.Matched)
			} else {
				totalDiff = totalDiff.Add(order.Qty)
			}
		}
		reason += ", draw: "
		_, err := user.ChangeBalance(reason, totalDiff)
		if err != nil {
			log.Error("user.ChangeBalance", reason, err)
		}
	}
}
