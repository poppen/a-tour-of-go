package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
