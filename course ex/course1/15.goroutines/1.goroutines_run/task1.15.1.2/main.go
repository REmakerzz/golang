package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)

	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")

}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	data := make(chan string)

	go func() {
		defer close(data)

		startTime := time.Now()
		for {
			select {
			case <-ticker.C:
				data <- message
			case <-time.After(d):
				return
			}
			if time.Since(startTime) >= d {
				return
			}
		}
	}()

	return data
}
