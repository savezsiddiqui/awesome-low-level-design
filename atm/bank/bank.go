package bank

import (
	"fmt"

	"example.com/atm/account"
	custom_error "example.com/atm/error"
)

type Bank struct {
	ifsccode             string
	name                 string
	branch               string
	atmCardToAccountsMap map[string]*account.Account // key is card number:pin
	accounts             []*account.Account
}

type CreateNewBank struct {
	IFSCCode string
	Name     string
	Branch   string
}

type LinkAtmCardToBankAccount struct {
	CardNumber string
	Pin        string
	Account    *account.Account
}

func NewBank(req *CreateNewBank) *Bank {
	return &Bank{
		ifsccode:             req.IFSCCode,
		name:                 req.Name,
		branch:               req.Branch,
		atmCardToAccountsMap: make(map[string]*account.Account),
		accounts:             []*account.Account{},
	}
}

func (b *Bank) AuthenticateCard(cardNumber string, pin string) (*account.Account, error) {
	key := fmt.Sprintf("%s:%s", cardNumber, pin)
	acc, ok := b.atmCardToAccountsMap[key]
	if !ok {
		return nil, custom_error.NewIncorrectCardOrPinError()
	}
	return acc, nil
}

func (b *Bank) AddAccount(acc *account.Account) {
	b.accounts = append(b.accounts, acc)
}

func (b *Bank) IssueNewAtmCard(req *LinkAtmCardToBankAccount) error {
	key := fmt.Sprintf("%s:%s", req.CardNumber, req.Pin)

	if _, exists := b.atmCardToAccountsMap[key]; exists {
		return custom_error.NewCardAlreadyExistsError()
	}

	b.atmCardToAccountsMap[key] = req.Account
	return nil
}
