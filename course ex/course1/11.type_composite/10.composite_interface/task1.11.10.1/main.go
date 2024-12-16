package main

import (
	"fmt"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprint("int")
	case string:
		return fmt.Sprint("string")
	case []int:
		return fmt.Sprint("[]int")
	case float64:
		return fmt.Sprint("float64")
	case float32:
		return fmt.Sprint("float32")
	default:
		return fmt.Sprint("Пустой интерфейс")
	}
}

func main() {
	var i interface{} = 42

	fmt.Println(getType(i))

}
