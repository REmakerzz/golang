package main

import (
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	const alphabet = "ABCDEFGHJKLMNOPQRSTUVWXYZ0123456789"
	var keySlice []string

	length := 4

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		key := make([]byte, length)
		for j := range key {
			key[j] = alphabet[rand.Intn(len(alphabet))]
		}
		keySlice = append(keySlice, string(key))
	}
	return strings.Join(keySlice, "-")
}

func main() {
}
