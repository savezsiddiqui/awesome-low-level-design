package atm

import (
	"time"

	"example.com/atm/bank"
	"example.com/atm/cash_dispenser"
	custom_error "example.com/atm/error"
	"example.com/atm/transaction"
	"github.com/google/uuid"
)

type Atm struct {
	bankingService *bank.Bank
	cashDispenser  *cash_dispenser.CashDispenser
}

func NewAtm(bankService *bank.Bank, cashDispenser *cash_dispenser.CashDispenser) *Atm {
	return &Atm{
		bankingService: bankService,
		cashDispenser:  cashDispenser,
	}
}

type AtmCardAuthenticationRequest struct {
	CardNumber string
	Pin        string
}

func (atm *Atm) GetBalance(req *AtmCardAuthenticationRequest) (float64, error) {
	acc, err := atm.bankingService.AuthenticateCard(req.CardNumber, req.Pin)
	if err != nil {
		return 0.0, err
	}
	return acc.GetBalance(), nil
}

func (atm *Atm) Withdraw(req *AtmCardAuthenticationRequest, amount float64) error {
	acc, err := atm.bankingService.AuthenticateCard(req.CardNumber, req.Pin)
	if err != nil {
		return err
	}

	if !atm.cashDispenser.HasSufficientCash(amount) {
		return custom_error.NewInsufficientCashError()
	}

	txn := transaction.Transaction{
		TransactionID:   uuid.NewString(),
		Timestamp:       time.Now(),
		TransactionType: transaction.Withdrawal,
		Amount:          amount,
	}

	if err := acc.Withdraw(amount); err != nil {
		return err
	}

	if err := atm.cashDispenser.Withdraw(amount); err != nil {
		return err
	}
	acc.AddTransaction(&txn)
	return nil
}

func (atm *Atm) Deposit(req *AtmCardAuthenticationRequest, amount float64) error {
	acc, err := atm.bankingService.AuthenticateCard(req.CardNumber, req.Pin)
	if err != nil {
		return err
	}

	txn := transaction.Transaction{
		TransactionID:   uuid.NewString(),
		Timestamp:       time.Now(),
		TransactionType: transaction.Deposit,
		Amount:          amount,
	}

	acc.Deposit(amount)
	atm.cashDispenser.Deposit(amount)
	acc.AddTransaction(&txn)
	return nil
}
