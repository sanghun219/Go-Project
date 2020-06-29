package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("we have no money")

// NewAccount Cunstructor
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit func
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance Get
func (a *Account) Balance() int {
	return a.balance
}

// WithDraw what
func (a *Account) WithDraw(amount int) error {
	if a.balance-amount < 0 {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(owner string) {
	a.owner = owner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\n Has ", a.Balance())
}
