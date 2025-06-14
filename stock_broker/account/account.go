package account

import (
	"sync"

	custom_error "example.com/stock_broker/error"
	"example.com/stock_broker/portfolio"
	"example.com/stock_broker/transaction"
	"github.com/google/uuid"
)

type Account struct {
	accountId          string
	balance            float64
	Portfolio          portfolio.Portfolio
	mu                 sync.RWMutex
	transactionHistory []transaction.Transaction
}

func (acc *Account) Deposit(money float64) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += money
}

func (acc *Account) Withdraw(money float64) error {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	if money > acc.balance {
		err := custom_error.InsufficentBalanceError{}
		return err
	}

	acc.balance -= money
	return nil
}

func (acc *Account) GetBalance() float64 {
	acc.mu.RLock()
	defer acc.mu.RUnlock()
	return acc.balance
}

func (acc *Account) GetAccountId() string {
	return acc.accountId
}

func (acc *Account) AddTransaction(newTransaction transaction.Transaction) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.transactionHistory = append(acc.transactionHistory, newTransaction)
}

func (acc *Account) GetTransactionHistory() []*transaction.Transaction {
	acc.mu.RLock()
	defer acc.mu.RUnlock()

	history := make([]*transaction.Transaction, len(acc.transactionHistory))
	for i := range acc.transactionHistory {
		history[i] = &acc.transactionHistory[i]
	}
	return history
}

func CreateNewAccount() *Account {
	return &Account{
		accountId: uuid.New().String(),
		Portfolio: portfolio.NewPortfolio(),
	}
}
