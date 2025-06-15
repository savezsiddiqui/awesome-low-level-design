package account

import (
	"sync"

	custom_error "example.com/atm/error"
	"example.com/atm/transaction"
)

type Account struct {
	name               string
	number             string
	balance            float64
	mu                 sync.RWMutex
	transactionHistory []*transaction.Transaction
}

type CreateAccountRequest struct {
	Name    string
	Number  string
	Balance float64
}

func NewAccount(req *CreateAccountRequest) *Account {
	return &Account{
		name:    req.Name,
		number:  req.Number,
		balance: req.Balance,
	}
}

func (acc *Account) GetName() string {
	return acc.name
}

func (acc *Account) GetNumber() string {
	return acc.number
}

func (acc *Account) GetBalance() float64 {
	acc.mu.RLock()
	defer acc.mu.RUnlock()
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if amount > acc.balance {
		return custom_error.NewInsufficientBalanceError()
	}

	acc.balance -= amount
	return nil
}

func (acc *Account) AddTransaction(txn *transaction.Transaction) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.transactionHistory = append(acc.transactionHistory, txn)
}

func (acc *Account) ShowTransactionHistory() []*transaction.Transaction {
	return acc.transactionHistory
}
