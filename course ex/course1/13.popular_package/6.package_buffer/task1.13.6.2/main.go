package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Hello\n,\n World!")

	buffer := bytes.NewBuffer(data)

	scanner := getScanner(buffer)

	if scanner == nil {
		panic("Expected non-nil reader, got nil")
	}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func getScanner(b *bytes.Buffer) *bufio.Scanner {
	return bufio.NewScanner(b)
}
