package main

import (
	"fmt"
)

var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} = func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

var Concat func(xs ...interface{}) interface{} = func(xs ...interface{}) interface{} {
	result := ""
	for _, x := range xs {
		if str, ok := x.(string); ok {
			result += str
		}
	}
	return result
}

var Sum func(xs ...interface{}) interface{} = func(xs ...interface{}) interface{} {
	sumInt := 0

	sumFloat := 0.0
	for _, x := range xs {
		switch v := x.(type) {
		case int:
			sumInt += v
		case float64:
			sumFloat += float64(v)
		}
	}
	if sumInt != 0 {
		return sumInt
	} else {
		return sumFloat
	}
}

func main() {
	fmt.Println(Operate(Concat, "Hello", " ", "World!"))
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5))
}
