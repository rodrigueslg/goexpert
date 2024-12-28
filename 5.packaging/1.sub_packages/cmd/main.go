package main

import (
	"fmt"

	"github.com/rodrigueslg/fc-goexpert/packaging/1/math"
)

func main() {
	m := math.Math{A: 3, B: 4}
	fmt.Println(m.Add())
}
