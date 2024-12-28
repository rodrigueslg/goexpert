package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c *Cliente) andou() {
	c.Nome = "Luis Gustavo"
	fmt.Printf("O cliente %v andou\n", c.Nome)
}

func main() {
	luis := Cliente{
		Nome: "Luis",
	}
	luis.andou()
	fmt.Printf("O valor da struct com nome é %v\n", luis.Nome)
}