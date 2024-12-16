package main

import "fmt"

func main() {
	intArr := [8]int{5, 2, 8, 1, 9, 3, 7, 4}
	floatArr := [8]float64{5.5, 2.2, 8.8, 1.1, 9.9, 3.3, 7.7, 4.4}

	sortedIntDesc := sortDescInt(intArr)
	sortedIntAsc := sortAscInt(intArr)
	sortedFloatDesc := sortDescFloat(floatArr)
	sortedFloatAsc := sortAscFloat(floatArr)

	fmt.Println("Sorted Int Array (Descending):", sortedIntDesc)
	fmt.Println("Sorted Int Array (Ascending):", sortedIntAsc)
	fmt.Println("Sorted Float Array (Descending):", sortedFloatDesc)
	fmt.Println("Sorted Float Array (Ascending):", sortedFloatAsc)
}

func sortDescInt(intArr [8]int) [8]int {
	n := len(intArr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if intArr[j] < intArr[j+1] {
				intArr[j], intArr[j+1] = intArr[j+1], intArr[j]
			}
		}
	}
	return intArr
}

func sortAscInt(intArr [8]int) [8]int {
	n := len(intArr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if intArr[j] > intArr[j+1] {
				intArr[j], intArr[j+1] = intArr[j+1], intArr[j]
			}
		}
	}
	return intArr
}

func sortDescFloat(floatArr [8]float64) [8]float64 {
	n := len(floatArr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if floatArr[j] < floatArr[j+1] {
				floatArr[j], floatArr[j+1] = floatArr[j+1], floatArr[j]
			}
		}
	}
	return floatArr
}

func sortAscFloat(floatArr [8]float64) [8]float64 {
	n := len(floatArr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if floatArr[j] > floatArr[j+1] {
				floatArr[j], floatArr[j+1] = floatArr[j+1], floatArr[j]
			}
		}
	}
	return floatArr
}
