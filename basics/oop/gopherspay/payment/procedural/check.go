package procedural

type CheckingAccount struct {
	AccountOwner string
	RoutingNumber string
	AccountNumber string
	Balance float32
}

func PayWithCheck(account *CheckingAccount, amount float32) bool {
	if amount > account.Balance {
		return false
	}
	account.Balance -= amount
	return true
}