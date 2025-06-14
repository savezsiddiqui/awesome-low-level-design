package atm

import (
	"errors"
	"fmt"

	"example.com/atm/bank"
	"example.com/atm/cash_dispenser"
	"example.com/atm/transaction"
)

type Atm struct {
	BankingService *bank.Bank
	CashDispenser  *cash_dispenser.CashDispenser
}

func (atm *Atm) GetBalance(accountNumber string) (float64, error) {
	if acc, ok := atm.BankingService.AccountMap[accountNumber]; ok {
		return acc.GetBalance(), nil
	}

	return 0.0, errors.New("no account exists with this number")
}

func (atm *Atm) Withdraw(accountNumber string, amount float64) (transaction.Transaction, error) {
	if acc, ok := atm.BankingService.AccountMap[accountNumber]; ok {
		if err := atm.CashDispenser.Withdraw(amount); err != nil {
			fmt.Println(err)
			return &transaction.WithdrawalTransaction{}, err
		}
		if err := transaction.NewWithdrawalTransaction(acc, amount).Execute(); err != nil {
			return 
		}
	}

	return &transaction.WithdrawalTransaction{}, errors.New("no account exists with this number")
}
