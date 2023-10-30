package main

import (
	"bank_api/internal"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	bank := internal.CreateBank()

	card := internal.CreateCard()
	account := internal.CreateAccount()
	account.AddCard(card)
	user := internal.CreateUser()
	user.AddAccount(account)
	bank.AddUser(user)
	account.Balance = 1000
	fmt.Println(user.ID)
	fmt.Println("Acc", account.ID)
	fmt.Println("card", card.ID)

	router := mux.NewRouter()
	router.Handle("/account/{UserID}/{AccountID}", &bank)
	router.Handle("/account/{UserID}/{AccountID}/{CardID}", &bank)
	router.Handle("/account/delete", &bank)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
