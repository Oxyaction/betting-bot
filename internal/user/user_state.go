package user

import (
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

type UserState struct {
	ContextRoute  string
	PreviousRoute string
	Match         string
	Coeff         decimal.Decimal
	Side          ob.Side
	Qty           decimal.Decimal
}
