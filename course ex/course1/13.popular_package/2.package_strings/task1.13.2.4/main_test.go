package main

import "testing"

func TestGenerateActivationKey(t *testing.T) {
	key := generateActivationKey()

	if len(key) != 19 {
		t.Errorf("Expected key lenght of 19, got %d", len(key))
	}

	expectedFormat := "XXXX-XXXX-XXXX-XXXX"
	if !isValidFormat(key, expectedFormat) {
		t.Errorf("Key format is invalid: %s", key)
	}
}

func isValidFormat(key, format string) bool {
	if len(key) != len(format) {
		return false
	}

	for i := 0; i < len(format); i++ {
		if format[i] == 'X' && !(key[i] >= 'A' && key[i] <= 'Z' || key[i] >= '0' && key[i] <= '9') {
			return false
		} else if format[i] != 'X' && format[i] != key[i] {
			return false
		}
	}

	return true
}
