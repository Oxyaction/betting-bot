package user

import "github.com/shopspring/decimal"

type History []BalanceChange

type BalanceChange struct {
	Reason string
	Diff   decimal.Decimal
	Actual decimal.Decimal
}
