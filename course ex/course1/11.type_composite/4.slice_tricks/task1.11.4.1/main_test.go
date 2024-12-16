package main

import (
	"bytes"
	"os"
	"testing"
)

func TestCut(t *testing.T) {
	testCases := []struct {
		input    []int
		start    int
		end      int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 3, []int{2, 3, 4}},
		{[]int{}, 0, 0, []int{}},
		{[]int{3, 2, 1, 5, 4}, -1, 2, []int{}},
		{[]int{3, 2, 1, 5, 4}, 2, -1, []int{}},
		{[]int{3, 2, 1, 5, 4}, 5, 0, []int{}},
		{[]int{3, 2, 1, 5, 4}, 0, 5, []int{}},
	}

	for _, tc := range testCases {
		result := Cut(tc.input, tc.start, tc.end)
		if len(result) != len(tc.expected) {
			t.Errorf("got %v, want: %v", result, tc.expected)
		}
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("got: %v, want: %v", result, tc.expected)
			}
		}
	}
}

func testMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	w.Close()

	os.Stdout = old
	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "[]\n"
	if stdout.String() != expected {
		t.Errorf("got: %v, want: %v", stdout.String(), expected)
	}
}
