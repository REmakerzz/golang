package main

import (
	"testing"
)

type testData struct {
	n        int
	expected int
}

func TestFibonacci(t *testing.T) {
	testCases := []testData{
		{n: 5, expected: 5},
		{n: 6, expected: 8},
		{n: 7, expected: 13},
		{n: 8, expected: 21},
		{n: 9, expected: 34},
		{n: 10, expected: 55},
	}

	for _, tc := range testCases {
		result := Fibonacci(tc.n)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %d, Expected: %d, Got: %d", tc, tc.expected, result)
		}
	}
}
