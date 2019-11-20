package main

import (
	"fmt"
	"math"
)

type verterx struct {
	X, Y float64
}

func (v verterx) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *verterx) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := verterx{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
