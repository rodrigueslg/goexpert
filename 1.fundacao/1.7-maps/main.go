package main

import "fmt"

func main() {
	salarios := map[string]int{"Luis": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Luis"])

	delete(salarios, "João")

	salarios["Luis"] = 6666
	fmt.Println(salarios["Luis"])

	sal := make(map[string]int)
	sal1 := map[string]int{}	
}