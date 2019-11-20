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

func main() {
	v := vertex{3, 4}
	fmt.Println(abs(v))
}
