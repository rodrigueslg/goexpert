package main

import "sync"

// thread 1 - main
func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)

	go publish(ch)
	go reader(ch, &wg)

	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		println(x)
		wg.Done()
	}
}
