package main

import "fmt"

type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

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

type Empresa struct {
	Nome	string
	Ativo	bool
}

func (e Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("A empresa %s foi desativada \n", e.Nome)
}

func main() {
	luis := Cliente{
		Nome: "Luis",
		Idade: 33,
		Ativo: true,		
	}
	
	empresa := Empresa{
		Nome: "Empresa teste",
		Ativo: true,
	}
	
	Desativacao(luis)
	Desativacao(empresa)
}