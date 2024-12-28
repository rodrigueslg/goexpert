package main

import "fmt"

type Endereco struct {
	Logradouro	string
	Numero		int
	Cidade		string
	Estado		string
}

type Cliente struct {
	Nome		string
	Idade		int
	Ativo		bool
	Endereco
}

func main() {
	luis := Cliente{
		Nome: "Luis",
		Idade: 33,
		Ativo: true,		
	}
	luis.Cidade = "Belo Horizonte"

	fmt.Printf("A cidade do cliente é %s \n", luis.Cidade)
}