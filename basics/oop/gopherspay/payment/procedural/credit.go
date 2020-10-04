package procedural

type CreditCard struct {
	OwnerName string
	CardNumber string
	ExpirationMonth int
	ExpirationYear int
	SecurityCode int
	AvailableCredit float32
}

func PayWithCredit(card *CreditCard, amount float32) bool {
	if card.AvailableCredit < amount {
		return false
	}
	card.AvailableCredit -= amount
	return true
}
