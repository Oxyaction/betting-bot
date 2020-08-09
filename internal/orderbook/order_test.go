package orderbook

import (
	"testing"

	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

func TestOrderLay(t *testing.T) {
	o := NewOrder()
	o.Side = ob.Lay
	o.Qty = decimal.NewFromInt(100)
	o.Coeff = decimal.NewFromInt(100)

}
