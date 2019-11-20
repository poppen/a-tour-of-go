package main

import (
	"fmt"
)

type vertex struct {
	X, Y float64
}

func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := vertex{3, 4}
	v.Scale(2)
	scaleFunc(&v, 10)

	p := &vertex{4, 3}
	p.Scale(3)
	scaleFunc(p, 8)

	fmt.Println(v, p)
}
