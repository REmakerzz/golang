package main

import (
	"bytes"
	"os"
	"testing"
)

func TestGetSubSlice(t *testing.T) {
	testCase := []struct {
		input    []int
		start    int
		end      int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, 6, []int{3, 4, 5, 6}},
		{[]int{2, 3, 1, 4, 2, 5, 2, 6, 7}, 1, 3, []int{3, 1}},
	}

	for _, tc := range testCase {
		result := getSubSlice(tc.input, tc.start, tc.end)
		if len(result) != len(tc.expected) {
			t.Errorf("Unxpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("Unxpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
			}
		}
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

	expected := "[3 4 5 6]\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
