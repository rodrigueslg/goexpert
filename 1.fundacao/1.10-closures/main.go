package main

import "fmt"

func main() {

	total := func() int {
		return sum(42, 88)
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {	
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}