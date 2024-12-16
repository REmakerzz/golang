package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func WithCustomerID(customerID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = customerID
	}
}

func WithItems(items []string) OrderOption {
	return func(o *Order) {
		o.Items = items
	}
}

func WithOrderDate(orderDate time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = orderDate
	}
}

func NewOrder(ID int, options ...OrderOption) *Order {
	o := &Order{
		ID: ID,
	}

	for _, option := range options {
		option(o)
	}
	return o
}

func main() {
	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems([]string{"item1", "item2"}),
		WithOrderDate(time.Now()))
	fmt.Printf("Order: %+v\n", order)
}
