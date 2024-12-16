package main

import (
	"fmt"
	"github.com/icrowley/fake"
)

func main() {
	fmt.Println(GenerateFakeData())
}

// GenerateFakeData return fake data
func GenerateFakeData() string {
	name := "Name: " + fake.FirstName()
	address := "Address: " + fake.StreetAddress()
	phone := "Phone: " + fake.Phone()
	email := "Email: " + fake.EmailAddress()
	return name + "\n" + address + "\n" + phone + "\n" + email
}
