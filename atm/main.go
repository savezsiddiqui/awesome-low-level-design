package main

import (
	"fmt"

	"example.com/atm/account"
	"example.com/atm/bank"
	"example.com/atm/card"
)

func main() {
	account1 := account.Account{
		Name:    "Savez Siddiqui",
		Number:  "123456789",
		Balance: 10000.00,
	}

	account2 := account.Account{
		Name:    "Ramesh Tiwari",
		Number:  "121234567",
		Balance: 9165.00,
	}

	bank := bank.Bank{
		AccountMap: map[string]*account.Account{
			account1.Number: &account1,
			account2.Number: &account2,
		},
	}

	if acc, ok := bank.AccountMap["123456789"]; ok {
		fmt.Println(acc.GetBalance())
		acc.Deposit(5000)
		fmt.Println(acc.GetBalance())
	}

	// adding a card
	card := card.Card{
		NameOnCard: account1.Name,
		CardNumber: account1.Number,
	}
	card.SetPin("7467")
	account1.AddCard(&card)

	if err := account1.Cards[0].Authenticate("7467"); err != nil {
		fmt.Println(err.Error())
	}

	account1.RemoveCard(&card)
	fmt.Println(len(account1.Cards))
}
