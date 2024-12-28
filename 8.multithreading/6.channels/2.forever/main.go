package main

// thread 1 - main
func main() {
	forever := make(chan bool)

	go func() {
		// thread 2
		for i := 0; i < 10; i++ {
			println(i)
		}

		forever <- true
	}()

	<-forever
}
