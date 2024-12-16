package main

import (
	"reflect"
	"testing"
)

func generateSlice(size int) []float64 {
	slice := make([]float64, size)
	for i := 0; i < size; i++ {
		slice[i] = float64(i)
	}
	return slice
}

func TestAverage(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{}, 0},
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{-1, -0, 1}, 0},
		{generateSlice(5), 2},
	}

	for _, test := range tests {
		result := average(test.input)
		if result != test.expected {
			t.Errorf("average(%v) = %f; expected %f", test.input, result, test.expected)
		}
	}

	slice1 := generateSlice(10)
	slice2 := generateSlice(10)

	if !reflect.DeepEqual(slice1, slice2) {
		t.Errorf("generateSlice(10) produced different results : %v and %v", slice1, slice2)
	}
}
