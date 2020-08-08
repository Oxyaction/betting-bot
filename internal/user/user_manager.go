package user

import (
	"fmt"

	ob "github.com/miktwon/orderbook"
	"github.com/shopspring/decimal"
)

type UserManager struct {
	users      map[int]*User
	userStates map[int]*UserState
}

func NewUserManager() *UserManager {
	return &UserManager{
		make(map[int]*User),
		make(map[int]*UserState),
	}
}

func (um *UserManager) Add(id int) {
	if _, ok := um.users[id]; ok {
		return
	}
	um.users[id] = NewUser(id)
	um.userStates[id] = &UserState{}
}

func (um *UserManager) GetState(id int) *UserState {
	return um.userStates[id]
}

func (um *UserManager) GetPreviousRoute(id int) string {
	if state, ok := um.userStates[id]; ok {
		return state.PreviousRoute
	}
	return "start"
}

func (um *UserManager) GetContextRoute(id int) string {
	if state, ok := um.userStates[id]; ok {
		return state.ContextRoute
	}
	return ""
}

func (um *UserManager) SetContextRoute(id int, route string) error {
	if _, ok := um.userStates[id]; !ok {
		return fmt.Errorf("UserState does not exist for user #%d", id)
	}
	um.userStates[id].ContextRoute = route
	return nil
}

func (um *UserManager) SetPreviousRoute(id int, route string) error {
	if _, ok := um.userStates[id]; !ok {
		return fmt.Errorf("UserState does not exist for user #%d", id)
	}
	um.userStates[id].PreviousRoute = route
	return nil
}

func (um *UserManager) SetCoeff(id int, coeff decimal.Decimal) error {
	if _, ok := um.userStates[id]; !ok {
		return fmt.Errorf("UserState does not exist for user #%d", id)
	}
	um.userStates[id].Coeff = coeff
	return nil
}

func (um *UserManager) SetQty(id int, qty decimal.Decimal) error {
	if _, ok := um.userStates[id]; !ok {
		return fmt.Errorf("UserState does not exist for user #%d", id)
	}
	um.userStates[id].Qty = qty
	return nil
}

func (um *UserManager) SetSide(id int, side ob.Side) error {
	if _, ok := um.userStates[id]; !ok {
		return fmt.Errorf("UserState does not exist for user #%d", id)
	}
	um.userStates[id].Side = side
	return nil
}

func (um *UserManager) SetMatch(id int, match string) error {
	if _, ok := um.userStates[id]; !ok {
		return fmt.Errorf("UserState does not exist for user #%d", id)
	}
	um.userStates[id].Match = match
	return nil
}
