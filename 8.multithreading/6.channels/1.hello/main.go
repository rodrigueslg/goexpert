package main

import "fmt"

// thread 1 - main
func main() {
	// canal criado na thread 1
	channel := make(chan string)

	msg := "thread 1 set this value"
	fmt.Println(msg)

	// thread 2 - go routine
	go func() {
		// thread 2 adiciona valor ao canal
		channel <- "thread 2 set this value"
	}()

	// thread 1 lê valor adicionado ao canal pela thread 2
	msg = <-channel
	fmt.Println(msg)
}
