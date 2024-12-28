package main

const a = "Hello, World!"

var (
	b bool = true
	c int = 10
	d string = "luis"
	e float64 = 1.2
	f ID = 1
)

type ID int

func main() {
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	
	z := "z"
	println(z)
}