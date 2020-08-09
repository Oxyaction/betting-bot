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

	u1 := user.NewUser(1)
	u1.ChangeBalance("test money", decimal.NewFromInt(100))
	o1 := orderbook.NewOrder()
	o1.Side = ob.Back
	o1.Coeff = decimal.NewFromFloat(1.5)
	o1.Qty = decimal.NewFromFloat(8)

	_, err := e.PlaceOrder(u1, o1)
	if err != nil {
		t.Fatal(err)
	}

	u2 := user.NewUser(2)
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

	t.Logf("balance1 %s", u1.GetBalance())
	t.Logf("balance2 %s", u2.GetBalance())

	err = em.SettleEvent(e.ID, WinSideBack)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("balance1 %s", u1.GetBalance())
	t.Logf("balance2 %s", u2.GetBalance())
}

func TestManager2(t *testing.T) {
	log := logrus.New()
	em := NewManager(log)
	e := em.NewEvent("eventCategory", "eventName", time.Now().Add(2*time.Second))

	u1 := user.NewUser(1)
	u1.ChangeBalance("test money", decimal.NewFromInt(100))
	o1 := orderbook.NewOrder()
	o1.Side = ob.Back
	o1.Coeff = decimal.NewFromFloat(10)
	o1.Qty = decimal.NewFromFloat(100)
	t.Logf("balance1 %s", u1.GetBalance())

	_, err := e.PlaceOrder(u1, o1)
	if err != nil {
		t.Fatal(err)
	}

	u2 := user.NewUser(2)
	u2.ChangeBalance("test money", decimal.NewFromInt(901))
	o2 := orderbook.NewOrder()
	o2.Side = ob.Lay
	o2.Coeff = decimal.NewFromFloat(10)
	o2.Qty = decimal.NewFromFloat(100)
	t.Logf("balance2 %s", u2.GetBalance())

	changedMap, err := e.PlaceOrder(u2, o2)
	if err != nil {
		t.Fatal(err)
	}

	for u, orders := range changedMap {
		fmt.Printf("User: %d:\n", u.ID())
		for i, o := range orders {
			fmt.Printf("  #%d S:%s Q:%s M:%s U:%s St:%s\n", i+1, o.Side, o.Qty, o.Matched, o.Unmatched, o.Status)
		}
	}

	t.Logf("balance1 %s", u1.GetBalance())
	t.Logf("balance2 %s", u2.GetBalance())

	err = em.SettleEvent(e.ID, WinSideBack)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("balance1 %s", u1.GetBalance())
	t.Logf("balance2 %s", u2.GetBalance())
}
