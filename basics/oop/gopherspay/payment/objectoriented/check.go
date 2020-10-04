package objectoriented

// All the fields are package protected and not accessible outside directly
type CheckingAccount struct {
	accountOwner string
	routingNumber string
	accountNumber string
	balance float32
}

// Factory function to create Checking account
func CreateCheckingAccount(accountOwner, routingNumber, accountNumber string) *CheckingAccount {
	return &CheckingAccount{
		accountOwner: accountOwner,
		routingNumber: routingNumber,
		accountNumber: accountNumber,
		balance: 250, // This should come from call to web service
	}
}

func (c CheckingAccount) ProcessPayment(amount float32) bool {
	return true
}

func (c CheckingAccount) Balance() float32 {
	return c.balance
}


