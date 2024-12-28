package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d got %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	workerCount := 10000

	// initialize workes
	for i := 0; i < workerCount; i++ {
		go worker(i, data)
	}

	// feed channel
	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
