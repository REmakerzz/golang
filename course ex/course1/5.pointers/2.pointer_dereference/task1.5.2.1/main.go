package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := Dereference(&a)
	d := Sum(&b, &c)
	fmt.Println(c)
	fmt.Println(d)
}
// Dereference разыменовывает указатель на int
func Dereference(n *int) int {
	return *n
}
// Sum складывает две переменные по значению
func Sum(a, b *int) int {
	return *a + *b
}
