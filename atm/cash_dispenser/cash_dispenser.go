package cash_dispenser

import (
	custom_error "example.com/atm/error"
)

type CreateCashDispenserRequest struct {
	InitialBalance float64
}

type CashDispenser struct {
	balance float64
}

func NewCashDispenser(req *CreateCashDispenserRequest) *CashDispenser {
	return &CashDispenser{
		balance: req.InitialBalance,
	}
}

func (cd *CashDispenser) Withdraw(amount float64) error {
	if amount > cd.balance {
		return custom_error.NewInsufficientCashError()
	}

	cd.balance -= amount
	return nil
}

func (cd *CashDispenser) GetBalance() float64 {
	return cd.balance
}

func (cd *CashDispenser) Deposit(amount float64) {
	cd.balance += amount
}

func (cd *CashDispenser) HasSufficientCash(amount float64) bool {
	return cd.balance >= amount
}
