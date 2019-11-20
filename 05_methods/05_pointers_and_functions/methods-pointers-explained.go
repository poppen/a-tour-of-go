package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := vertex{3, 4}
	scale(&v, 10)
	fmt.Println(abs(v))
}
