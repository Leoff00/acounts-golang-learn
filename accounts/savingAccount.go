package accounts

import "github.com/acounts-golang-learn/clients"

type SavingAccount struct {
	Holder                                 clients.Holder
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso."
	} else {
		return "Saque realizado com sucesso."
	}
}

func (c *SavingAccount) Deposit(depositValue float64) (string, float64) {

	if depositValue > 0 {
		c.balance += depositValue
		return "Saque realizado com sucesso.", c.balance
	} else {
		return "Valor do deposito menor que zero.", c.balance
	}

}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
