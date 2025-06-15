package main

import (
	"fmt"

	"example.com/atm/account"
	"example.com/atm/atm"
	"example.com/atm/bank"
	"example.com/atm/cash_dispenser"
)

func main() {
	// 1. Create a new bank instance
	myBank := bank.NewBank(&bank.CreateNewBank{
		IFSCCode: "BANK123",
		Name:     "My Awesome Bank",
		Branch:   "Main Branch",
	})
	fmt.Println("Bank created:", myBank)

	// 2. Create a new cashDispenser instance
	cashDispenser := cash_dispenser.NewCashDispenser(&cash_dispenser.CreateCashDispenserRequest{
		InitialBalance: 10000.00,
	})
	fmt.Println("Cash Dispenser created with initial balance:", cashDispenser.GetBalance())

	// 3. Create a new Atm instance
	myATM := atm.NewAtm(myBank, cashDispenser)
	fmt.Println("ATM instance created.")

	// 4. Add some accounts into the bank
	account1 := account.NewAccount(&account.CreateAccountRequest{
		Name:    "Alice Smith",
		Number:  "ACC001",
		Balance: 1500.00,
	})
	myBank.AddAccount(account1)
	fmt.Println("Account 1 added:", account1.GetName(), "Balance:", account1.GetBalance())

	account2 := account.NewAccount(&account.CreateAccountRequest{
		Name:    "Bob Johnson",
		Number:  "ACC002",
		Balance: 500.00,
	})
	myBank.AddAccount(account2)
	fmt.Println("Account 2 added:", account2.GetName(), "Balance:", account2.GetBalance())

	account3 := account.NewAccount(&account.CreateAccountRequest{
		Name:    "Charlie Brown",
		Number:  "ACC003",
		Balance: 20000.00,
	})
	myBank.AddAccount(account3)
	fmt.Println("Account 3 added:", account3.GetName(), "Balance:", account3.GetBalance())

	// 5. Issue some atm cards, from multiple accounts
	card1 := &bank.LinkAtmCardToBankAccount{CardNumber: "1111222233334444", Pin: "1234", Account: account1}
	myBank.IssueNewAtmCard(card1)
	fmt.Println("Card issued for Alice Smith.")

	card2 := &bank.LinkAtmCardToBankAccount{CardNumber: "5555666677778888", Pin: "5678", Account: account2}
	myBank.IssueNewAtmCard(card2)
	fmt.Println("Card issued for Bob Johnson.")

	card3 := &bank.LinkAtmCardToBankAccount{CardNumber: "9999888877776666", Pin: "9012", Account: account1}
	myBank.IssueNewAtmCard(card3)
	fmt.Println("Second card issued for Alice Smith (different card number/pin).")

	card4 := &bank.LinkAtmCardToBankAccount{CardNumber: "0000111122223333", Pin: "3456", Account: account3}
	myBank.IssueNewAtmCard(card4)
	fmt.Println("Card issued for Charlie Brown.")

	fmt.Println("\n--- Transaction Tests ---")

	// Test 1: Successful Deposit
	fmt.Println("\n--- Test 1: Successful Deposit (Alice Smith) ---")
	depositReq1 := &atm.AtmCardAuthenticationRequest{CardNumber: "1111222233334444", Pin: "1234"}
	depositAmount1 := 200.00
	err := myATM.Deposit(depositReq1, depositAmount1)
	if err != nil {
		fmt.Println("Deposit failed:", err)
	} else {
		fmt.Printf("Successfully deposited %.2f into Alice's account. New balance: %.2f\n", depositAmount1, account1.GetBalance())
	}

	// Test 2: Successful Withdrawal
	fmt.Println("\n--- Test 2: Successful Withdrawal (Alice Smith) ---")
	withdrawReq1 := &atm.AtmCardAuthenticationRequest{CardNumber: "1111222233334444", Pin: "1234"}
	withdrawAmount1 := 300.00
	err = myATM.Withdraw(withdrawReq1, withdrawAmount1)
	if err != nil {
		fmt.Println("Withdrawal failed:", err)
	} else {
		fmt.Printf("Successfully withdrew %.2f from Alice's account. New balance: %.2f\n", withdrawAmount1, account1.GetBalance())
	}

	// Test 3: Invalid Authentication Attempt (Wrong PIN)
	fmt.Println("\n--- Test 3: Invalid Authentication Attempt (Wrong PIN) ---")
	invalidAuthReq := &atm.AtmCardAuthenticationRequest{CardNumber: "1111222233334444", Pin: "9999"}
	err = myATM.Withdraw(invalidAuthReq, 50.00)
	if err != nil {
		fmt.Println("Withdrawal with wrong PIN failed as expected:", err)
	} else {
		fmt.Println("Error: Withdrawal unexpectedly succeeded with wrong PIN.")
	}

	// Test 4: Withdrawing more than balance in account
	fmt.Println("\n--- Test 4: Withdrawing more than balance (Bob Johnson) ---")
	withdrawReq2 := &atm.AtmCardAuthenticationRequest{CardNumber: "5555666677778888", Pin: "5678"}
	withdrawAmount2 := 1000.00 // Bob only has 500
	err = myATM.Withdraw(withdrawReq2, withdrawAmount2)
	if err != nil {
		fmt.Println("Withdrawal more than balance failed as expected:", err)
	} else {
		fmt.Println("Error: Withdrawal unexpectedly succeeded with insufficient balance.")
	}
	fmt.Printf("Bob's balance after failed attempt: %.2f\n", account2.GetBalance())

	// Test 5: Withdrawing more than cash in cash_dispenser (Charlie Brown) ---
	fmt.Println("\n--- Test 5: Withdrawing more than cash in dispenser (Charlie Brown) ---")
	withdrawReq3 := &atm.AtmCardAuthenticationRequest{CardNumber: "0000111122223333", Pin: "3456"}

	// Charlie withdraws a large amount to drain the dispenser
	drainAmount := 9800.00
	err = myATM.Withdraw(withdrawReq3, drainAmount)
	if err != nil {
		fmt.Println("Initial large withdrawal failed:", err)
	} else {
		fmt.Printf("Successfully withdrew %.2f from Charlie's account to drain dispenser. New Charlie balance: %.2f\n", drainAmount, account3.GetBalance())
	}
	fmt.Printf("Cash dispenser balance after draining: %.2f\n", cashDispenser.GetBalance())

	// Charlie attempts to withdraw more than remaining cash in dispenser
	insufficientCashWithdrawal := 200.00 // Dispenser has 100.00 left, this should fail
	err = myATM.Withdraw(withdrawReq3, insufficientCashWithdrawal)
	if err != nil {
		fmt.Println("Withdrawal more than cash in dispenser failed as expected:", err)
	} else {
		fmt.Println("Error: Withdrawal unexpectedly succeeded with insufficient cash in dispenser.")
	}
	fmt.Printf("Charlie's balance after failed attempt: %.2f\n", account3.GetBalance())
	fmt.Printf("Cash dispenser balance after failed attempt: %.2f\n", cashDispenser.GetBalance())

	// Test 6: Another successful withdrawal (Bob Johnson)
	fmt.Println("\n--- Test 6: Another successful withdrawal (Bob Johnson) ---")
	depositAmount2 := 2000.00 // Deposit more for Bob
	_ = myATM.Deposit(withdrawReq2, depositAmount2)
	fmt.Printf("Bob's balance after deposit: %.2f\n", account2.GetBalance())

	withdrawAmount3 := 100.00
	err = myATM.Withdraw(withdrawReq2, withdrawAmount3)
	if err != nil {
		fmt.Println("Withdrawal failed:", err)
	} else {
		fmt.Printf("Successfully withdrew %.2f from Bob's account. New balance: %.2f\n", withdrawAmount3, account2.GetBalance())
	}

	// Print the transaction history
	fmt.Println("\n--- Transaction History ---")
	fmt.Println("\nAlice Smith's Transaction History:")
	for _, txn := range account1.ShowTransactionHistory() {
		fmt.Printf("  ID: %s, Type: %s, Amount: %.2f, Time: %s\n", txn.TransactionID, txn.TransactionType.String(), txn.Amount, txn.Timestamp.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nBob Johnson's Transaction History:")
	for _, txn := range account2.ShowTransactionHistory() {
		fmt.Printf("  ID: %s, Type: %s, Amount: %.2f, Time: %s\n", txn.TransactionID, txn.TransactionType.String(), txn.Amount, txn.Timestamp.Format("2006-01-02 15:04:05"))
	}

	fmt.Println("\nCharlie Brown's Transaction History:")
	for _, txn := range account3.ShowTransactionHistory() {
		fmt.Printf("  ID: %s, Type: %s, Amount: %.2f, Time: %s\n", txn.TransactionID, txn.TransactionType.String(), txn.Amount, txn.Timestamp.Format("2006-01-02 15:04:05"))
	}

	fmt.Printf("\nFinal Cash Dispenser Balance: %.2f\n", cashDispenser.GetBalance())
}
