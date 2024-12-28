package main

import "fmt"

// write only channel must have <- on right side
func escrever(text string, ch chan<- string) {
	ch <- text
}

// read only channel must have <- on left side
func ler(ch <-chan string) {
	fmt.Println(<-ch)
}

func main() {
	ch := make(chan string, 1)
	go escrever("Olá", ch)
	ler(ch)
}
