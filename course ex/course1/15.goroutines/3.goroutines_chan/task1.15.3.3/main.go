package main

import (
	"fmt"
	"sync"
	"time"
)

type sema chan struct {
}

func New(n int) sema {
	return make(sema, n)
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s sema) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := New(n)

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sem.Inc(1)
			time.Sleep(1 * time.Second)
			sem.Dec(1)
			fmt.Println(n)
		}(num)

		wg.Wait()
	}
}
