package main

import "fmt"

func main() {
	var x int = 10
	var s string = "Hello"
	var y float64 = 2.4
	var k bool = true

	changeInt(&x)
	changeFloat(&y)
	changeString(&s)
	changeBool(&k)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s)
	fmt.Println(k)
}

func changeInt(a *int) {
	*a = 20
}

func changeFloat(b *float64) {
	*b = 6.28
}

func changeString(c *string) {
	*c = "Goodbye, world!"
}

func changeBool(d *bool) {
	*d = false
}
