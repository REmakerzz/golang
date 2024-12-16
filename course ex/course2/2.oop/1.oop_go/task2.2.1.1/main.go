package main

import (
	"errors"
	"fmt"
)

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	balance float64
}

type SavingsAccount struct {
	balance float64
}

func (c *CurrentAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	c.balance += amount
	return nil
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > c.balance {
		return errors.New("Not enough balance")
	}
	c.balance -= amount
	return nil
}

func (c *CurrentAccount) Balance() float64 {
	return c.balance
}

func (s *SavingsAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	s.balance += amount
	return nil
}

func (s *SavingsAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if s.balance <= 500 {
		return errors.New("balance is less than 500")
	}
	s.balance -= amount
	return nil
}

func (s *SavingsAccount) Balance() float64 {
	return s.balance
}

func ProcessAccount(a Account) {
	err := a.Deposit(500)
	if err != nil {
		fmt.Println("Error depositing:", err)
	}
	err = a.Withdraw(200)
	if err != nil {
		fmt.Println("Error withdrawing:", err)
	}
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{}

	ProcessAccount(c)
	ProcessAccount(s)
}
