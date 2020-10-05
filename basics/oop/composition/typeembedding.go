package main

import "fmt"

type Account struct {}

func (acc *Account) AvailableFunds() float32 {
	fmt.Println("Listing funds")
	return 0
}

func (acc *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

type CreditAccount struct {
	// Embedding the Account type into the CreditAccount
	// Any calls that are not of CreditAccount will automatically be delegated to Account
	Account
}


// Resolving Conflicts
type CredAccount struct {}

func (acc *CredAccount) AvailableFunds() float32 {
	fmt.Println("Getting credit funds")
	return 0
}

type CheckingAccount struct {}

func (acc *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds")
	return 0
}

type HybridAccount struct {
	CredAccount
	CheckingAccount
}

// Resolving the ambiguous call by providing implementation ourselves
func (acc *HybridAccount) AvailableFunds() float32 {
	return acc.CredAccount.AvailableFunds() + acc.CheckingAccount.AvailableFunds()
}

func main() {
	ca := &CreditAccount{}
	ca.AvailableFunds()

	ha := &HybridAccount{}
	// This will give compilation error because of ambiguous call if Hybrid Account doesn't implement the method itself
	// and help the compiler in resolving the conflict
	ha.AvailableFunds()
}