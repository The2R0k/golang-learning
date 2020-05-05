package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	y := x
	for prev := 0.0; prev != y; {
		prev = y
		y = (y + x/y) / 2
	}
	return y
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
