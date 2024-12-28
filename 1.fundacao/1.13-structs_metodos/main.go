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

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado \n", c.Nome)
}

func main() {
	luis := Cliente{
		Nome: "Luis",
		Idade: 33,
		Ativo: true,		
	}
	luis.Desativar()
}