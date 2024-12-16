package main

import (
	"fmt"
	"time"
)

func timeout(timeout time.Duration) func() bool {
	start := time.Now()
	done := make(chan bool)

	go func() {
		time.Sleep(timeout)
		done <- false
	}()

	return func() bool {
		select {
		case result := <-done:
			return result
		default:
			return time.Since(start) < timeout
		}
	}
}

func main() {
	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}
