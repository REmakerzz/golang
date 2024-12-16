package main

import (
	"bytes"
	"os"
	"testing"
)

func TestAppendInt(t *testing.T) {
	testCases := []struct {
		input    []int
		a        int
		b        int
		c        int
		expected []int
	}{
		{[]int{1, 2, 3}, 4, 5, 6, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 2, 2}, 3, 3, 3, []int{2, 2, 2, 3, 3, 3}},
	}

	for _, tc := range testCases {
		result := appendInt(tc.input, tc.a, tc.b, tc.c)
		if len(result) != len(tc.expected) {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("got %v, want: %v", result, tc.expected)
			}
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

	expected := "[1 2 3] [1 2 3 4 5 6]\n"
	if stdout.String() != expected {
		t.Errorf("got %v, want: %v", stdout.String(), expected)
	}
}
