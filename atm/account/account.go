package account

import (
	"errors"
	"sync"

	"example.com/atm/card"
)

type Account struct {
	Name    string
	Number  string
	Balance float64
	mu      sync.RWMutex
	Cards   []*card.Card
}

func (acc *Account) GetBalance() float64 {
	acc.mu.RLock()
	defer acc.mu.RUnlock()
	return acc.Balance
}

func (acc *Account) Deposit(amount float64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.Balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if amount > acc.Balance {
		return errors.New("this amount is greater than Balance")
	}

	acc.Balance -= amount
	return nil
}

func (acc *Account) AddCard(card *card.Card) {
	acc.Cards = append(acc.Cards, card)
}

func (acc *Account) RemoveCard(input *card.Card) {
	var newCards []*card.Card

	for _, c := range acc.Cards {
		if c != input {
			newCards = append(newCards, c)
		}
	}

	acc.Cards = newCards
}
