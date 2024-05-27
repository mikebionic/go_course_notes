package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdraw", amount, "\b, balance is now", b.balance)
	}
}

// command pattern
type Command interface {
	Call()
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account: account, action: action, amount: amount}
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
	case Withdraw:
		b.account.Withdraw(b.amount)
	}
}

func main() {
	ba := BankAccount{}
	cmd := NewBankAccountCommand(&ba, Deposit, 100)
	cmd.Call()
	fmt.Println(ba)
	cmd2 := NewBankAccountCommand(&ba, Withdraw, 50)
	cmd2.Call()
	fmt.Println(ba)
}
