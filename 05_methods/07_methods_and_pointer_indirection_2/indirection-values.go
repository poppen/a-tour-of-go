package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func absFunc(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(absFunc(v))

	p := vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(absFunc(p))
}
