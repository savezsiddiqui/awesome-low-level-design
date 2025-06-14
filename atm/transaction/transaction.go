package transaction

import (
	"time"

	"example.com/atm/account"
)

const (
	Deposit = iota
	Withdrawal
)

type Transaction interface {
	Execute() error
}

type BaseTransaction struct {
	Account         *account.Account
	TransactionType int
	Amount          float64
	TimeStamp       time.Time
	TransactionID   string
}

type WithdrawalTransaction struct {
	BaseTransaction
}

func NewWithdrawalTransaction(account *account.Account, amount float64) Transaction {
	return &WithdrawalTransaction{
		BaseTransaction: BaseTransaction{
			Account:         account,
			TransactionType: Withdrawal,
			Amount:          amount,
			TimeStamp:       time.Now().UTC(),
			TransactionID:   "",
		},
	}
}

func (w *WithdrawalTransaction) Execute() error {
	if err := w.Account.Withdraw(w.Amount); err != nil {
		return err
	}

	return nil
}

type DepositTransaction struct {
	BaseTransaction
}

func (d *DepositTransaction) Execute() error {
	d.Account.Deposit(d.Amount)
	return nil
}

func NewDepositTransaction(account *account.Account, amount float64) Transaction {
	return &DepositTransaction{
		BaseTransaction: BaseTransaction{
			Account:         account,
			TransactionType: Deposit,
			Amount:          amount,
			TimeStamp:       time.Now().UTC(),
			TransactionID:   "",
		},
	}
}
