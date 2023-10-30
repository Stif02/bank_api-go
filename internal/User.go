package internal

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Firstname string
	Lastname  string
	Accounts  map[uuid.UUID]*Account
}

func CreateUser() *User {
	return &User{
		ID:       uuid.Must(uuid.NewRandom()),
		Accounts: make(map[uuid.UUID]*Account),
	}
}

func (u *User) AddAccount(a *Account) {
	u.Accounts[a.ID] = a
}
