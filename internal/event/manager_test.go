package event

import (
	"fmt"
	"testing"
	"time"

	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gitlab.com/fireferretsbet/tg-bot/internal/orderbook"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

func TestManager(t *testing.T) {
	log := logrus.New()
	em := NewManager(log)
	e := em.NewEvent("eventCategory", "eventName", time.Now().Add(2*time.Second))

	u1 := user.NewUser()
	u1.ChangeBalance("test money", decimal.NewFromInt(100))
	o1 := orderbook.NewOrder()
	o1.Side = ob.Back
	o1.Coeff = decimal.NewFromFloat(1.5)
	o1.Qty = decimal.NewFromFloat(8)

	_, err := e.PlaceOrder(u1, o1)
	if err != nil {
		t.Fatal(err)
	}

	u2 := user.NewUser()
	u2.ChangeBalance("test money", decimal.NewFromInt(100))
	o2 := orderbook.NewOrder()
	o2.Side = ob.Lay
	o2.Coeff = decimal.NewFromFloat(1.5)
	o2.Qty = decimal.NewFromFloat(1)

	changedMap, err := e.PlaceOrder(u2, o2)
	if err != nil {
		t.Fatal(err)
	}

	for u, orders := range changedMap {
		fmt.Println(u.ID(), orders)
	}

	err = em.SettleEvent(e.ID, WinSideLay)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("balance1 %s", u1.GetBalance())
	t.Logf("balance2 %s", u2.GetBalance())
}
