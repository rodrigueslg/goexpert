package main

func main() {
	for i:=0; i<5; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}

	i := 0
	for i < 5 {
		println(i)
		i++
	}

	for {
		println("Hello, World!")
	}
}