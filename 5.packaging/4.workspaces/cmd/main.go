package main

import (
	"github.com/google/uuid"
	"github.com/rodrigueslg/fc-goexpert/packaging/3/math"
)

func main() {
	m := &math.Math{3, 4}
	println(m.Add())

	println(uuid.New().String())
}
