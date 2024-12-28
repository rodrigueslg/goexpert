package main

import (
	"fmt"
	"github.com/rodrigueslg/curso-go-fc/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Printf("Resultado: %v\n", s)
}
