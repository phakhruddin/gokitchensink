package main

import "fmt"

// Defining the structure for an account
type Account struct {
	ID      int
	Balance float64
}

// Implement the function to transfer funds between accounts
func transferFunds(amount float64, source *Account, destination *Account) {
	// Write your implementation here
	fmt.Println("amount to xfer", amount, "source.Balance:", source.Balance, "destination.Balance:", destination.Balance)

	if amount <= source.Balance {
		fmt.Println("XFER STARTED:amount to xfer", amount, "source.Balance:", source.Balance, "destination.Balance:", destination.Balance)
		destination.Balance = amount + destination.Balance
		source.Balance = source.Balance - amount
		fmt.Println("source.Balance:", source.Balance, "destination.Balance:", destination.Balance)
		goto transactionComplete
	} else {
		goto insufficientBalance
	}
transactionComplete:
	fmt.Println("transactionComplete")
	return
insufficientBalance:
	fmt.Println("Insufficient fund")
}

func main() {
	// Creating two accounts with initial balances
	account1 := Account{ID: 1, Balance: 100.0}
	account2 := Account{ID: 2, Balance: 50.0}

	// Demonstrating various transactions
	transferFunds(50.0, &account1, &account2)
	transferFunds(200.0, &account1, &account2)
}
