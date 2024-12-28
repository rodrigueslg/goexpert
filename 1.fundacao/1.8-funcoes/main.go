package main

import "fmt"

func main() {
	 result, bool := sum(20, 100)
	 fmt.Println(result)
	 fmt.Println(bool)
}

func sum(a, b int) (int, bool) {
	result := a + b
	if result >= 50 {
		return result, true
	}
	return result, false
}