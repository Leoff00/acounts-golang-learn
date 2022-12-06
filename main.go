package main

import (
	"fmt"

	"github.com/acounts-golang-learn/accounts"
)

func PaymentTicket(account verifyAccount, ticketValue float64) {
	account.Withdraw(ticketValue)
}

type verifyAccount interface {
	Withdraw(valor float64) string
}

func main() {
	acc1 := accounts.SavingAccount{}
	acc1.Deposit(10000)
	acc1.Withdraw(200)
	PaymentTicket(&acc1, 60)
	cash := acc1.GetBalance()

	fmt.Println(cash)
}
