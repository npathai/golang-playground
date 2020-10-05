package main

import (
	"../../payment/objectoriented"
	"fmt"
	"strings"
)

func main() {
	const amount = 500

	fmt.Println("Paying with cash")
	cash := objectoriented.CreateCashAccount()
	cash.ProcessPayment(amount)

	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	card := objectoriented.CreateCreditCard(
		"Narendra Pathai",
		"1111-2222-3333-4444",
		5,
		2021,
		123)

	fmt.Println("Paying with card card")
	fmt.Printf("Initial balance: %.2f\n", card.AvailableCredit())
	card.ProcessPayment(amount)
	fmt.Printf("Balance now: %.2f\n", card.AvailableCredit())
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")

	checking := objectoriented.CreateCheckingAccount(
		"Narendra Pathai",
		"01001001",
		"123456789",
	)

	fmt.Println("Paying with check")
	fmt.Printf("Initial balance: %.2f\n", checking.Balance())
	checking.ProcessPayment(amount)
	fmt.Printf("Balance now: %.2f\n", checking.Balance())

	fmt.Println("Hmm, not enough balance in account. Nothing we can do")
	fmt.Printf(strings.Repeat("*", 10) + "\n\n")
}
