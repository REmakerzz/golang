package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)

	expected := "map[amet:3 lorem:1 sit:3]\n"
	if stdout.String() != expected {
		t.Errorf("got: %q, want: %q", stdout.String(), expected)
	}
}

func TestCountWordsInText(t *testing.T) {
	testCases := []struct {
		inputTxt   string
		inputWords []string
		expected   map[string]int
	}{
		{
			inputTxt: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris. 
        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor. 
        Praesent et diam eget libero egestas mattis sit amet vitae augue.`,
			inputWords: []string{"sit", "amet", "lorem"},
			expected:   map[string]int{"amet": 2, "lorem": 1, "sit": 3},
		},
		{
			inputTxt:   "",
			inputWords: []string{},
			expected:   map[string]int{},
		},
		{
			inputTxt:   "text",
			inputWords: []string{},
			expected:   map[string]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.inputTxt, func(t *testing.T) {
			got := CountWordsInText(tc.inputTxt, tc.inputWords)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("CountWordsInText() = %v, want %v", got, tc.expected)
			}
		})
	}
}
