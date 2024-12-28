package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// thread 1 - main
func main() {
	// thread 2 - goroutine 1
	go task("A")
	// thread 3 - goroutine 2
	go task("B")

	// thread 4 - goroutine 3
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(15 * time.Second)
}
