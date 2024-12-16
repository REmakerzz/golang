package main

import "fmt"

func getVariableType(variable interface{}) string {
	switch variable.(type) {
	case int:
		return fmt.Sprintf("%T", variable)
	case string:
		return fmt.Sprintf("%T", variable)
	case float64:
		return fmt.Sprintf("%T", variable)
	default:
		return "Error variable"
	}
}

func main() {
	var num int = 10

	var str string = "Hello"
	var num2 float64 = 3.14

	fmt.Println(getVariableType(num))
	fmt.Println(getVariableType(str))
	fmt.Println(getVariableType(num2))
}
