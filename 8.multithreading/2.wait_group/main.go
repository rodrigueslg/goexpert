package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// thread 1 - main
func main() {
	wg := sync.WaitGroup{}
	wg.Add(25)

	// thread 2 - goroutine 1
	go task("A", &wg)
	// thread 3 - goroutine 2
	go task("B", &wg)

	// thread 4 - goroutine 3
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			wg.Done()
		}
	}()

	wg.Wait()
}
