package accounts

import "github.com/acounts-golang-learn/clients"

type CurrentAccount struct {
	Holder                      clients.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canCashout := withdrawValue > 0 && withdrawValue <= c.balance

	if canCashout {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso."
	} else {
		return "Saque realizado com sucesso."
	}
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {

	if depositValue > 0 {
		c.balance += depositValue
		return "Saque realizado com sucesso.", c.balance
	} else {
		return "Valor do deposito menor que zero.", c.balance
	}

}

func (c *CurrentAccount) Transfer(transferValue float64, destinyAccount *CurrentAccount) bool {
	if transferValue < c.balance {
		c.balance -= transferValue
		destinyAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
