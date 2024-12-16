package main

import (
	"bytes"
	"os"
	"testing"
)

type testData struct {
	input    [8]int
	expected int
}

func TestSum(t *testing.T) {
	testCases := []testData{
		{[8]int{2, 4, 6, 4, 1, 8, 5, 6}, 36},
		{[8]int{9, 9, 9, 9, 9, 9, 9, 9}, 72},
		{[8]int{1, 1, 1, 1, 1, 1, 1, 1}, 8},
		{[8]int{0, 0, 0, 0, 0, 0, 0, 0}, 0},
	}
	for _, tc := range testCases {
		result := sum(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestAverage(t *testing.T) {
	testCases := []struct {
		input    [8]int
		expected float64
	}{
		{[8]int{1, 2, 3, 4, 5, 6, 7, 8}, 4.5},
		{[8]int{2, 2, 2, 2, 2, 2, 2, 2}, 2},
	}

	for _, tc := range testCases {
		result := average(tc.input)
		if result != float64(tc.expected) {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestAverageFloat(t *testing.T) {
	testCases := []struct {
		input    [8]float64
		expected float64
	}{
		{[8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}, 5},
		{[8]float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}, 0},
	}

	for _, tc := range testCases {
		result := averageFloat(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input    [8]int
		expected [8]int
	}{
		{[8]int{1, 2, 3, 4, 5, 6, 7, 8}, [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
		{[8]int{8, 7, 6, 5, 4, 3, 2, 1}, [8]int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tc := range testCases {
		result := reverse(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
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

	expected := "36\n4.5\n5\n[8 7 6 5 4 3 2 1]\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}
