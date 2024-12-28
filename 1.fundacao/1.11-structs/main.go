package main

import "fmt"

type Cliente struct {
	Nome string
	Idade int
	Ativo bool
}

func main() {
	luis := Cliente{
		Nome: "Luis",
		Idade: 33,
		Ativo: true,
	}	

	fmt.Printf("O nome do cliente é %s \n", luis.Nome)
	
	luis.Nome = "Luis Gustavo"
	fmt.Printf("O nome do cliente é %s \n", luis.Nome)
}