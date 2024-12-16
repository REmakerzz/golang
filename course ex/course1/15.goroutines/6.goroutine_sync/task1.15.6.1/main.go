package main

import (
	"fmt"
	"sync"
)

func waitGroupExample(goroutines ...func() string) string {
	var wg sync.WaitGroup

	results := make(chan string, len(goroutines))

	for _, goroutine := range goroutines {
		wg.Add(1)
		go func(g func() string) {
			defer wg.Done()
			results <- g()
		}(goroutine)
	}

	wg.Wait()
	close(results)

	var output string
	for result := range results {
		output += result + "\n"
	}
	return output
}

func main() {
	count := 2
	goroutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func() string {
			return fmt.Sprintf("goroutine %d", j)
		}
	}

	fmt.Println(waitGroupExample(goroutines...))
}
