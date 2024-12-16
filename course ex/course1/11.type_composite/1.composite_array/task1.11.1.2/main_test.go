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

	expected := "Sorted Int Array (Descending): [9 8 7 5 4 3 2 1]\nSorted Int Array (Ascending): [1 2 3 4 5 7 8 9]\nSorted Float Array (Descending): [9.9 8.8 7.7 5.5 4.4 3.3 2.2 1.1]\nSorted Float Array (Ascending): [1.1 2.2 3.3 4.4 5.5 7.7 8.8 9.9]\n"
	if stdout.String() != expected {
		t.Errorf("got %q, want %q", stdout.String(), expected)
	}
}

func TestSortDescInt(t *testing.T) {
	testCases := []struct {
		input    [8]int
		expected [8]int
	}{
		{[8]int{5, 4, 2, 7, 1, 8, 3, 9}, [8]int{9, 8, 7, 5, 4, 3, 2, 1}},
		{[8]int{1, 2, 3, 4, 5, 7, 8, 9}, [8]int{9, 8, 7, 5, 4, 3, 2, 1}},
	}

	for _, tc := range testCases {
		result := sortDescInt(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestSortAscInt(t *testing.T) {
	testCases := []struct {
		input    [8]int
		expected [8]int
	}{
		{[8]int{9, 8, 7, 5, 4, 3, 2, 1}, [8]int{1, 2, 3, 4, 5, 7, 8, 9}},
		{[8]int{1, 3, 5, 2, 6, 9, 8, 4}, [8]int{1, 2, 3, 4, 5, 6, 8, 9}},
	}

	for _, tc := range testCases {
		result := sortAscInt(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestSortDescFloat(t *testing.T) {
	testCases := []struct {
		input    [8]float64
		expected [8]float64
	}{
		{[8]float64{1.1, 3.3, 8.8, 4.4, 9.9, 6.6, 2.2, 7.7}, [8]float64{9.9, 8.8, 7.7, 6.6, 4.4, 3.3, 2.2, 1.1}},
		{[8]float64{2.2, 8.8, 3.3, 9.9, 4.4, 7.7, 5.5, 6.6}, [8]float64{9.9, 8.8, 7.7, 6.6, 5.5, 4.4, 3.3, 2.2}},
	}

	for _, tc := range testCases {
		result := sortDescFloat(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}

func TestSortAscFloat(t *testing.T) {
	testCases := []struct {
		input    [8]float64
		expected [8]float64
	}{
		{[8]float64{1.1, 3.3, 8.8, 4.4, 9.9, 6.6, 2.2, 7.7}, [8]float64{1.1, 2.2, 3.3, 4.4, 6.6, 7.7, 8.8, 9.9}},
		{[8]float64{2.2, 8.8, 3.3, 9.9, 4.4, 7.7, 5.5, 6.6}, [8]float64{2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}},
	}

	for _, tc := range testCases {
		result := sortAscFloat(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc, tc.expected, result)
		}
	}
}
