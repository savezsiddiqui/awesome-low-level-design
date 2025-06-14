package cash_dispenser

import "errors"

type CashDispenser struct {
	Balance float64
}

func (cd *CashDispenser) Withdraw(amount float64) error {
	if amount > cd.Balance {
		return errors.New("insufficient funds in cash dispenser")
	}

	cd.Balance -= amount
	return nil
}
