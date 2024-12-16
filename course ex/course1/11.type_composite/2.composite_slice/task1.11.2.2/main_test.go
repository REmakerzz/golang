package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "8\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want: %q", stdout.String(), expected)
	}
}

func TestMaxDifference(t *testing.T) {
	testCase := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1}, 0},
		{[]int{}, 0},
		{[]int{9, 3, 3, 8}, 6},
		{[]int{4, 3, 2, 1}, 3},
	}

	for _, tc := range testCase {
		result := MaxDifference(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}
