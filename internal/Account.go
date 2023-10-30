package internal

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID
	Active  bool
	Balance float64
	Limit   float64
	Cards   map[uuid.UUID]*Card
}

func (a *Account) CloseAccount() {
	a.Active = false
}

func (a *Account) SetLimit(requiredLimit float64) {
	a.Limit = requiredLimit
}

func (a *Account) ViewBalance() float64 {
	return a.Balance
}

func (a *Account) AddCard(c *Card) {
	a.Cards[c.ID] = c
}

func CreateAccount() *Account {
	account := &Account{
		ID:     uuid.Must(uuid.NewRandom()),
		Active: true,
		Cards:  make(map[uuid.UUID]*Card),
	}
	account.AddCard(CreateCard())
	return account
}
