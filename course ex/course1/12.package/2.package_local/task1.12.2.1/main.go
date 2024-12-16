package main

import (
	"fmt"

	mymath "studentgit.kata.academy/REmakerzz/go-kata/course1/12.package/2.package_local/task1.12.2.1/test"
)

func main() {
	x := 4.0

	y := 2.0

	fmt.Println("Square root of", x, "is", mymath.Sqrt(x))
	fmt.Println("Ceil of", x, "is", mymath.Ceil(x))
	fmt.Println("Floor of", x, "is", mymath.Floor(x))
	fmt.Println("Power of", x, "to", y, "is", mymath.Pow(x, y))
	fmt.Println("Max of", x, "and", y, "is", mymath.Max(x, y))
	fmt.Println("Min of", x, "and", y, "is", mymath.Min(x, y))
}
