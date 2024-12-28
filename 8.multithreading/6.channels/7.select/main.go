package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Microsecond)
		c2 <- 2
	}()

	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-c1:
			println("received", msg1)
		case msg2 := <-c2:
			println("received", msg2)
		case <-time.After(time.Second * 3):
			println("timeout")
		default:
			println("default")
		}

		time.Sleep(time.Microsecond)
	}
}
