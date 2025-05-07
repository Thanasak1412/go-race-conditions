package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	counter++
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go increment(&wg)
	}

	wg.Wait()

	fmt.Println("Final counter", counter)
}
