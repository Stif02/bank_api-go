package internal

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Bank struct {
	Users map[uuid.UUID]*User
}

func (b *Bank) AddUser(u *User) {
	b.Users[u.ID] = u
}

func CreateBank() Bank {
	bank := Bank{
		Users: make(map[uuid.UUID]*User),
	}
	bank.AddUser(CreateUser())
	return bank
}

func (b *Bank) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		vars := mux.Vars(r)

		if _, ok := vars["AccountID"]; ok {
			err := json.NewEncoder(w).Encode(b.Users[uuid.MustParse(vars["UserID"])].Accounts[uuid.MustParse(vars["AccountID"])].Balance)
			if err != nil {
				log.Fatal(err)
			}
		}
		if _, ok := vars["CardID"]; ok {
			err := json.NewEncoder(w).Encode(b.Users[uuid.MustParse(vars["UserID"])].Accounts[uuid.MustParse(vars["AccountID"])].Cards[uuid.MustParse(vars["CardID"])].ViewCardStatus())
			if err != nil {
				log.Fatal(err)
			}
		}
	case http.MethodPost:
		vars := mux.Vars(r)
		if vars["account"] == "create" {
			CreateAccount()
		}
	}
}
