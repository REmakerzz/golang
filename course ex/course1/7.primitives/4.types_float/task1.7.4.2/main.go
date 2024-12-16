package main

import (
	"fmt"
	"github.com/mattevans/dinero"
	"time"
)

func main() {

	rate := currencyPairRate("USD", "EUR", 100.0)
	fmt.Println(rate)
}

func currencyPairRate(curr1, curr2 string, cash float64) float64 {
	client := dinero.NewClient(
		"6c05caa8bd1840bdac58e4a90fa32b59",
		"USD",
		20*time.Minute,
	)
	a, err := client.Rates.Get(curr1)
	if err != nil {
		panic(err)
	}
	b, err := client.Rates.Get(curr2)
	if err != nil {
		panic(err)
	}
	result := (*a * cash) * *b
	return result
}