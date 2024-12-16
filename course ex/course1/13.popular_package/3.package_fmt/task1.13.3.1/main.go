package main

import "fmt"

func generateMathString(operands []int, operator string) string {
	if len(operands) == 0 {
		return "Zero operands"
	}

	expression := ""
	for i, operand := range operands {
		if i > 0 {
			expression += fmt.Sprintf(" %s ", operator)
			fmt.Println(i, expression)
		}
		expression += fmt.Sprintf("%d", operand)
	}

	result := operands[0]

	for _, operand := range operands[1:] {
		switch operator {
		case "+":
			result += operand
		case "-":
			result -= operand
		case "*":
			result *= operand
		case "/":
			if operand != 0 {
				result /= operand
			} else {
				return "Error: Div by zero"
			}
		default:
			return "Error operator"
		}
	}
	return fmt.Sprintf("%s = %d", expression, result)
}

func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "+"))
}
