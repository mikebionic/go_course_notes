package main

import "fmt"

// whenever you have a change in the system you can take a snapshot of the system

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

// for the initial memento
func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance: balance}, &Memento{balance}
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{b.balance}
}

func (b *BankAccount) Restore(m *Memento) {
	b.balance = m.Balance
}

func main() {
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	fmt.Println(ba)
	ba.Restore(m1)
	fmt.Println(ba)
	ba.Restore(m2)
	fmt.Println(ba)

	ba.Restore(m0)
	fmt.Println(ba)
	// &{175}
	// &{150}
	// &{175}
	// &{100}
}
