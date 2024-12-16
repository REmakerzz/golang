package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		now := time.Now()

		currentTime := now.Format("15:04:05")
		currentDate := now.Format("2006-01-02")

		fmt.Print("\033[H\033[2J")
		fmt.Printf("Текущее время: %s\n", currentTime)
		fmt.Printf("Текущая дата: %s\n", currentDate)

		time.Sleep(1 * time.Second)
	}
}
