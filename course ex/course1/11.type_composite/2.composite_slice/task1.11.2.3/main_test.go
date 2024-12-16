package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFindSingleNumber(t *testing.T) {

	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 4, 3, 2, 1}, 5},
		{[]int{7, 8, 7, 9, 8}, 9},
		{[]int{10}, 10},
		{[]int{1, 1, 2, 2, 3}, 3},
		{[]int{0, 1, 0}, 1},
	}

	for _, tc := range testCases {
		result := findSingleNumber(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "5\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
