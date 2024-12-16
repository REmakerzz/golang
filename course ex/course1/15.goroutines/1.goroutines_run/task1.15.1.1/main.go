package main

import (
	"fmt"
	"time"
)

func main() {

	stop := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Горутина завершила работу")
		stop <- true
	}()

	timer := time.NewTimer(5 * time.Second)

	data := NotifyOnTimer(timer, stop)

	for v := range data {
		fmt.Println(v)
	}

}

func NotifyOnTimer(timer *time.Timer, stop chan bool) <-chan string {
	data := make(chan string)

	go func() {
		for {
			select {
			case <-timer.C:
				data <- "Таймер сработал"
				timer.Stop()
				close(data)
				return
			case <-stop:
				data <- "Горутина завершила работу раньше, чем таймер сработал"
				timer.Stop()
				close(data)
				return
			}
		}
	}()

	return data
}
