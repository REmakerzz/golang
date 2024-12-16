package main

import (
	"bytes"
	"os"
	"testing"
)

type testData struct {
	n        int
	expected int
}

func TestFibonacciDP(t *testing.T) {
	testCases := []testData{
		{n: 1, expected: 1},
		{n: 5, expected: 5},
		{n: 6, expected: 8},
		{n: 7, expected: 13},
		{n: 8, expected: 21},
		{n: 9, expected: 34},
		{n: 10, expected: 55},
	}

	for _, tc := range testCases {
		result := FibonacciDP(tc.n)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %d, Expected: %d, Got: %d", tc, tc.expected, result)
		}
	}
}

func TestFibonacciBinet(t *testing.T) {
	testCases := []testData{
		{n: 1, expected: 1},
		{n: 5, expected: 5},
		{n: 6, expected: 8},
		{n: 7, expected: 13},
		{n: 8, expected: 21},
		{n: 9, expected: 34},
		{n: 10, expected: 55},
	}

	for _, tc := range testCases {
		result := FibonacciBinet(tc.n)
		if result != tc.expected {
			t.Errorf("Unexpected result, Input: %d, Expected: %d, Got: %d", tc, tc.expected, result)
		}
	}
}

func BenchmarkFibonacciBinet(b *testing.B) {
	number := 5
	for i := 0; i < b.N; i++ {
		FibonacciBinet(number)
	}
}

func BenchmarkFibonacciDP(b *testing.B) {
	number := 5
	for i := 0; i < b.N; i++ {
		FibonacciDP(number)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "5\n5\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
