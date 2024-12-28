package main

func main() {
	a := 10

	println(a)
	println(&a)

	var ponteiro *int = &a
	println(ponteiro)
	println(*ponteiro)

	*ponteiro = 20
	println(*ponteiro)

	b := &a
	println(b)
	println(*b)

	a = 30
	println(b)
	println(*b)
}