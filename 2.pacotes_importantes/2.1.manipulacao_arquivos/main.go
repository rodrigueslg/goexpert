package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// escrita
	tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	f.Close()
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", tamanho)	

	// leitura completa
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// leitura em partes
	arquivo2, err := os.Open("arquivo.txt")	
	if err != nil {
		panic(err)
	}
	
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 1)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// deletar
	os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}