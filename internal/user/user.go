package user

import (
	"errors"
	"sync"

	"github.com/shopspring/decimal"
)

var (
	ErrNotEnoughFounds = errors.New("user: not enough founds")
)

type User struct {
	mu      sync.RWMutex
	id      int
	balance decimal.Decimal
	history History
}

func NewUser(id int) *User {
	return &User{
		id: id,
	}
}

func (u *User) ID() int {
	return u.id
}

func (u *User) GetBalance() (b decimal.Decimal) {
	u.mu.RLock()
	b = u.balance
	u.mu.RUnlock()
	return
}

func (u *User) ChangeBalance(reason string, diff decimal.Decimal) (bc BalanceChange, err error) {
	u.mu.Lock()
	bc, err = u.changeBalance(reason, diff)
	u.mu.Unlock()
	return
}

func (u *User) changeBalance(reason string, diff decimal.Decimal) (bc BalanceChange, err error) {
	bc.Actual = u.balance.Add(diff)
	if bc.Actual.IsNegative() {
		err = ErrNotEnoughFounds
		return
	}
	bc.Diff = diff
	bc.Reason = reason

	u.balance = bc.Actual
	u.history = append(u.history, bc)
	return
}
