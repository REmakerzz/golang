package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	randomString := GenerateRandomString(-1)
	fmt.Println(randomString, len(randomString))
}

func GenerateRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	const alphabet = "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var b strings.Builder

	b.Grow(length)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		b.WriteByte(alphabet[randomIndex])
	}

	return b.String()
}
