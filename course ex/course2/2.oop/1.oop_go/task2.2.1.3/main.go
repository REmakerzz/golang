package main

import (
	"errors"
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity < 0 {
		return errors.New("quantity must be positive")
	}
	d.orderDetails[item] = quantity
	return nil
}

func (d *DineInOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; ok {
		delete(d.orderDetails, item)
		return nil
	}
	return errors.New("item not found")
}

func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity < 0 {
		return errors.New("quantity must be positive")
	}
	t.orderDetails[item] = quantity
	return nil
}

func (t *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := t.orderDetails[item]; ok {
		delete(t.orderDetails, item)
		return nil
	}
	return errors.New("item not found")
}

func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}

func ManageOrder(o Order) {
	err := o.AddItem("Pizza", 2)
	if err != nil {
		fmt.Println(err)
	}
	err = o.AddItem("Burger", 1)
	if err != nil {
		fmt.Println(err)
	}
	err = o.RemoveItem("Pizza")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
