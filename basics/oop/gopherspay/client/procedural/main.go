package main

import (
	"fmt"
	"strings"
)
import "../../payment/procedural"

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	procedural.PayWithCash(amount)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	card := &procedural.CreditCard{
		OwnerName:       "Narendra Pathai",
		CardNumber:      "1111-2222-3333-4444",
		ExpirationMonth: 5,
		ExpirationYear:  2021,
		SecurityCode:    123,
		AvailableCredit: 5000,
	}

	fmt.Println("Paying with card card")
	fmt.Printf("Initial balance: %.2f\n", card.AvailableCredit)
	procedural.PayWithCredit(card, amount)
	fmt.Printf("Balance now: %.2f\n", card.AvailableCredit)
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := &procedural.CheckingAccount{
		AccountOwner: "Narendra Pathai",
		RoutingNumber: "01001001",
		AccountNumber: "123456789",
		Balance: 250,
	}

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: %.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("Balance now: %.2f\n", checking.Balance)

	fmt.Println("Hmm, not enough balance in account. We can fix that!")
	fmt.Println("Changing account balance")
	checking.Balance = 5000

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: %.2f\n", checking.Balance)
	procedural.PayWithCheck(checking, amount)
	fmt.Printf("Balance now: %.2f\n", checking.Balance)

}
