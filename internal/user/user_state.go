package user

import (
	"github.com/google/uuid"
	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

type UserState struct {
	ContextRoute  string
	PreviousRoute string
	Event         uuid.UUID
	Coeff         decimal.Decimal
	Side          ob.Side
	Qty           decimal.Decimal
}
