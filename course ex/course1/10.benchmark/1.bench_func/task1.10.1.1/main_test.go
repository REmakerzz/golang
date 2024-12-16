package main

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
		{20, 6765},
	}

	for _, test := range tests {
		result := Fibonacci(test.input)
		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30)
	}
}
