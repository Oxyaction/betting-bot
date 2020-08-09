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
	ErrNotEnoughFounds       = errors.New("orderbook: not enough founds")
	ErrOrderNotFound         = errors.New("orderbook: order not found")
	ErrOrderInvalidOrderUser = errors.New("orderbook: invalid order-user")
	ErrOrderAlreadyMatched   = errors.New("orderbook: order already matched")
	ErrOrderAlreadyCanceled  = errors.New("orderbook: order already canceled")
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

var one = decimal.NewFromInt(1)

//func (o Order) Liability() decimal.Decimal {
//	if o.Side == ob.Back {
//		return o.Qty
//	}
//
//	return o.Coeff.Sub(one).Mul(o.Qty)
//}

func (o Order) Settle(winSide ob.Side) decimal.Decimal {
	if o.Side != winSide {
		return o.Unmatched
	}

	c := o.Coeff.Sub(one)
	if o.Side == ob.Back {
		//fmt.Println("------>", o.Side, o.Coeff, o.Matched, o.Matched.Mul(o.Coeff), o.Qty)
		return o.Matched.Mul(o.Coeff).Add(o.Unmatched)
	}

	//fmt.Println("------>", o.Side, o.Coeff, o.Matched, o.Matched.Div(c), o.Qty)
	return o.Matched.Div(c).Add(o.Unmatched)
}
