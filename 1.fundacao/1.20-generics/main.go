package main

func SomaInteiro(m map[string]int) int {
	var soma int 
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyFloat float64

type Number interface {
	int | ~float64
}

func SomaGeneric[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T comparable](a T, b T) bool {
	return a==b
}

func main() {
	m:= map[string]int{"Luis": 1000, "João": 2000, "Maria": 3000}
	println(SomaInteiro(m))
	println(SomaGeneric(m))

	m2:= map[string]float64{"Luis": 1000.10, "João": 2000.20, "Maria": 3000.30}
	println(SomaFloat(m2))
	println(SomaGeneric(m2))

	m3:= map[string]MyFloat{"Luis": 1000.10, "João": 2000.20, "Maria": 3000.30}
	println(SomaGeneric(m3))

	println(compara(10, 11))
}
