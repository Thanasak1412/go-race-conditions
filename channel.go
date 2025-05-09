package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	var readWg sync.WaitGroup

	readWg.Add(1)

	go func() {
		sum := 0
		for val := range ch {
			sum += val
		}
		fmt.Println("Final Counter:", sum)
		readWg.Done()
	}()

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			ch <- 1
			wg.Done()
		}()
	}

	wg.Wait()
	close(ch)
	readWg.Wait()
}
