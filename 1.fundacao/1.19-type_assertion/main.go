package main

import "fmt"

type x interface {}

func main() {
	var x interface{} = "Luis Gustavo"
	println(x.(string))

	res, ok := x.(int)
	fmt.Printf("o valor de res é %v e o resultado de ok é %v\n", res, ok)
}
