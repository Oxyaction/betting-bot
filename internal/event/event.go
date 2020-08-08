package event

import (
	"errors"
	"sync"
	"time"

	"github.com/coreos/pkg/progressutil"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	ob "gitlab.com/fireferretsbet/tg-bot/internal/orderbook"
	"gitlab.com/fireferretsbet/tg-bot/internal/user"
)

type WinSide string

const (
	WinSideBack = "back"
	WinSideLay  = "lay"
	WinSideDraw = "draw"
)

var (
	ErrAlreadySettled = errors.New("event: already settled")
	ErrMatchStarted   = errors.New("event: match started")
	ErrNotFound       = errors.New("event: not found")
)

type Event struct {
	ID       uuid.UUID
	Name     string
	Category string
	StartAt  time.Time
	ClosedAt time.Time
	WinSide  WinSide

	started bool
	settled bool
	mu      sync.RWMutex
	winLine *ob.OrderBook
}

func NewEvent(
	log *logrus.Logger,
	name string,
	category string,
	startAt time.Time,
) *Event {
	return &Event{
		ID:       uuid.New(),
		Name:     name,
		Category: category,
		StartAt:  startAt,

		winLine: ob.NewOrderBook(log),
	}
}

func (e *Event) Start() (err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.started {
		err = progressutil.ErrAlreadyStarted
		return
	}

	e.started = true

	return
}

func (e *Event) Settle(winSide WinSide) (err error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.settled {
		err = ErrAlreadySettled
		return
	}
	e.settled = true
	e.WinSide = winSide
	e.ClosedAt = time.Now()

	switch winSide {
	case WinSideBack:
		e.winLine.SettleBack(e.Name)
	case WinSideLay:
		e.winLine.SettleLay(e.Name)
	default:
		e.winLine.SettleDraw(e.Name)
	}

	return
}

func (e *Event) PlaceOrder(user *user.User, order ob.Order) (changedMap map[*user.User][]ob.Order, err error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.started {
		err = ErrMatchStarted
		return
	}
	if e.settled {
		err = ErrAlreadySettled
		return
	}

	return e.winLine.Place(user, order)
}

func (e *Event) CancelOrder(user *user.User, orderID uuid.UUID) (order ob.Order, err error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.started {
		err = progressutil.ErrAlreadyStarted
		return
	}
	if e.settled {
		err = ErrAlreadySettled
		return
	}

	e.winLine.SettleBack()

	return e.winLine.Cancel(user, orderID)
}
