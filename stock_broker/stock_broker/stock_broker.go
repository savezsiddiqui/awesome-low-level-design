package stock_broker

import (
	"sync"
	"time"

	"example.com/stock_broker/account"
	custom_error "example.com/stock_broker/error"
	"example.com/stock_broker/stock"
	"example.com/stock_broker/transaction"
)

type StockBroker struct {
	Stocks   map[string]*stock.Stock
	Accounts map[string]*account.Account
	mu       sync.Mutex
}

func (sb *StockBroker) BuyStock(accountId string, stock *stock.Stock, qty int) error {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	var err error

	account, ok := sb.Accounts[accountId]
	if !ok {
		err = custom_error.NewAccountNotRegisteredError()
		return err
	}

	if (stock.GetPrice() * float64(qty)) > account.GetBalance() {
		err = custom_error.NewInsufficentBalanceError()
		return err
	}

	account.Portfolio.AddStock(stock, qty)
	newTransaction := transaction.Transaction{
		Timestamp:                     time.Now(),
		Stock:                         *stock,
		StockPriceatTimeofTransaction: stock.GetPrice(),
		Qty:                           qty,
		OrderType:                     transaction.Buy,
	}
	account.AddTransaction(newTransaction)
	return nil
}

func (sb *StockBroker) SellStock(accountId string, stock *stock.Stock, qty int) error {
	sb.mu.Lock()
	defer sb.mu.Unlock()

	var err error

	account, ok := sb.Accounts[accountId]
	if !ok {
		err = custom_error.NewAccountNotRegisteredError()
		return err
	}

	err = account.Portfolio.RemoveStock(stock, qty)
	if err != nil {
		return err
	}

	account.Deposit(stock.GetPrice() * float64(qty))
	newTransaction := transaction.Transaction{
		Timestamp:                     time.Now(),
		Stock:                         *stock,
		StockPriceatTimeofTransaction: stock.GetPrice(),
		Qty:                           qty,
		OrderType:                     transaction.Sell,
	}
	account.AddTransaction(newTransaction)
	return nil
}
