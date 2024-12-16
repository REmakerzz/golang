package main

import "fmt"

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:

		return false
	}
}

func main() {

	channel := make(chan int)
	fmt.Println(trySend(channel, 1))
	fmt.Println(trySend(channel, 2))
	fmt.Println(trySend(channel, 3))
	close(channel)
}
