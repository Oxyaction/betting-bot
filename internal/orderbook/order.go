package orderbook

import (
	"errors"
	"github.com/google/uuid"
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

type OrderStatus string

const (
	OrderStatusNew      OrderStatus = "NEW"
	OrderStatusPlaced   OrderStatus = "PLACED"
	OrderStatusPartial  OrderStatus = "PARTIAL"
	OrderStatusMatched  OrderStatus = "MATCHED"
	OrderStatusCanceled OrderStatus = "CANCELED"
)

var (
	ErrNotEnoughFounds = errors.New("orderbook: not enough founds")
)

type OrderMap map[uuid.UUID]Order

type Order struct {
	ID        uuid.UUID
	Status    OrderStatus
	Side      ob.Side
	Coeff     decimal.Decimal
	Qty       decimal.Decimal
	Unmatched decimal.Decimal
	Matched   decimal.Decimal
}

func NewOrder() Order {
	return Order{
		ID:     uuid.New(),
		Status: OrderStatusNew,
	}
}
