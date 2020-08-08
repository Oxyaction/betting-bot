package event

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	ob "gitlab.com/fireferretsbet/tg-bot/internal/orderbook"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type WinSide string

const (
	WinSideBack = "back"
	WinSideLay  = "lay"
	WinSideDraw = "draw"
)

type Event struct {
	ID       uuid.UUID
	Name     string
	StartAt  time.Time
	ClosedAt time.Time
	WinSide  WinSide

	settled bool
	mu      sync.RWMutex
	winLine *ob.OrderBook
}

func NewEvent(name string, startAt time.Time) *Event {
	return &Event{
		ID:      uuid.New(),
		Name:    name,
		StartAt: startAt,

		winLine: ob.NewOrderBook(),
	}
}

func (e *Event) Settle(winSide WinSide) (err error) {
	e.mu.Lock()
	if e.settled {
		err = fmt.Errorf("event already settled")
		return
	}
	e.settled = true
	e.WinSide = winSide
	e.ClosedAt = time.Now()

	switch winSide {
	case WinSideBack:
		e.winLine.SettleBack()
	case WinSideLay:
		e.winLine.SettleLay()
	default:
		e.winLine.SettleDraw()
	}

	e.mu.Unlock()
	return
}

func (e *Event) Place(user *user.User, order ob.Order) (changedMap map[*user.User][]ob.Order, err error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.settled {
		err = fmt.Errorf("event already settled")
		return
	}

	return e.winLine.Place(user, order)
}

func (e *Event) Cancel(user *user.User, orderID uuid.UUID) (order ob.Order, err error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if e.settled {
		err = fmt.Errorf("event already settled")
		return
	}

	return e.winLine.Cancel(user, orderID)
}
