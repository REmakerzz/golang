package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type SavingsAccount struct {
	balance float64
	mu      sync.Mutex
}

type CheckingAccount struct {
	balance float64
	mu      sync.Mutex
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.balance += amount
}

func (s *SavingsAccount) Withdraw(amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.balance-amount < 1000 {
		return errors.New("not enough funds for withdrawal, minimum balance must be 1000")
	}

	s.balance -= amount
	return nil
}

func (s *SavingsAccount) Balance() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.balance
}

func (c *CheckingAccount) Deposit(amount float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.balance += amount
}

func (c *CheckingAccount) Withdraw(amount float64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.balance < amount {
		return errors.New("not enough funds for withdrawal")
	}

	c.balance -= amount
	return nil
}

func (c *CheckingAccount) Balance() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.balance
}

type Customer struct {
	ID      int
	Name    string
	Account Account
}

type CustomerOption func(*Customer)

func NewCustomer(id int, options ...CustomerOption) *Customer {
	c := &Customer{
		ID: id,
	}

	for _, option := range options {
		option(c)
	}
	return c
}

func WithName(name string) CustomerOption {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}

func main() {
	savings := &SavingsAccount{}
	savings.Deposit(1099)

	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	err := customer.Account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Customer: %v, Balance: %v\n", customer.Name, customer.Account.Balance())
}
