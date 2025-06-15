package transaction

import (
	"fmt"
	"time"
)

type TransactionType uint

const (
	Deposit TransactionType = iota
	Withdrawal
)

func (tt TransactionType) String() string {
	switch tt {
	case Deposit:
		return "Deposit"
	case Withdrawal:
		return "Withdrawal"
	default:
		return fmt.Sprintf("Unknown TransactionType (%d)", tt)
	}
}

type Transaction struct {
	TransactionID   string
	TransactionType TransactionType
	Amount          float64
	Timestamp       time.Time
}
