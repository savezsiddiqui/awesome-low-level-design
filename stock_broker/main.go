package main

import (
	"fmt"

	"example.com/stock_broker/account"
	"example.com/stock_broker/stock"
	"example.com/stock_broker/stock_broker"
)

func main() {
	fmt.Println("Starting Stock Broker Demo")

	// 1. Create a new instance of stock_broker
	sb := &stock_broker.StockBroker{
		Stocks:   make(map[string]*stock.Stock),
		Accounts: make(map[string]*account.Account),
	}

	// 2. Add some stocks
	appleStock := stock.NewStock("AAPL", 150.00)
	sb.Stocks["AAPL"] = appleStock

	googleStock := stock.NewStock("GOOG", 2500.00)
	sb.Stocks["GOOG"] = googleStock

	fmt.Println("\n--- Stocks Added ---")
	for _, s := range sb.Stocks {
		fmt.Printf("Stock: %s, Price: %.2f\n", s.GetName(), s.GetPrice())
	}

	// 3. Add 2 accounts
	acc1 := account.CreateNewAccount()
	acc1.Deposit(10000.00)
	sb.Accounts[acc1.GetAccountId()] = acc1
	fmt.Printf("\nAccount 1 created with ID: %s, Balance: %.2f\n", acc1.GetAccountId(), acc1.GetBalance())

	acc2 := account.CreateNewAccount()
	acc2.Deposit(5000.00)
	sb.Accounts[acc2.GetAccountId()] = acc2
	fmt.Printf("Account 2 created with ID: %s, Balance: %.2f\n", acc2.GetAccountId(), acc2.GetBalance())

	// 4. Purchase some stocks
	fmt.Println("\n--- Purchasing Stocks ---")
	err := sb.BuyStock(acc1.GetAccountId(), appleStock, 10)
	if err != nil {
		fmt.Printf("Error buying Apple stock for Account 1: %v\n", err)
	} else {
		fmt.Printf("Account 1 purchased 10 AAPL. New Balance: %.2f\n", acc1.GetBalance())
	}

	err = sb.BuyStock(acc2.GetAccountId(), googleStock, 2)
	if err != nil {
		fmt.Printf("Error buying Google stock for Account 2: %v\n", err)
	} else {
		fmt.Printf("Account 2 purchased 2 GOOG. New Balance: %.2f\n", acc2.GetBalance())
	}

	// 5. Sell some stocks
	fmt.Println("\n--- Selling Stocks ---")
	err = sb.SellStock(acc1.GetAccountId(), appleStock, 5)
	if err != nil {
		fmt.Printf("Error selling Apple stock for Account 1: %v\n", err)
	} else {
		fmt.Printf("Account 1 sold 5 AAPL. New Balance: %.2f\n", acc1.GetBalance())
	}

	err = sb.SellStock(acc2.GetAccountId(), googleStock, 1)
	if err != nil {
		fmt.Printf("Error selling Google stock for Account 2: %v\n", err)
	} else {
		fmt.Printf("Account 2 sold 1 GOOG. New Balance: %.2f\n", acc2.GetBalance())
	}

	// 6. Print transaction history
	fmt.Println("\n--- Transaction History for Account 1 ---")
	history1 := acc1.GetTransactionHistory()
	for _, t := range history1 {
		fmt.Printf("Timestamp: %s, Type: %s, Stock: %s, Price: %.2f, Quantity: %d\n",
			t.Timestamp.Format("2006-01-02 15:04:05"), t.OrderType, t.Stock.GetName(), t.StockPriceatTimeofTransaction, t.Qty)
	}

	fmt.Println("\n--- Transaction History for Account 2 ---")
	history2 := acc2.GetTransactionHistory()
	for _, t := range history2 {
		fmt.Printf("Timestamp: %s, Type: %s, Stock: %s, Price: %.2f, Quantity: %d\n",
			t.Timestamp.Format("2006-01-02 15:04:05"), t.OrderType, t.Stock.GetName(), t.StockPriceatTimeofTransaction, t.Qty)
	}

	fmt.Println("\nStock Broker Demo Completed")
}
