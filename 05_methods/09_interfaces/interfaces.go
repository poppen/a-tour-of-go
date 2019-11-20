package main

import (
	"fmt"
	"math"
)

type abser interface {
	Abs() float64
}

func main() {
	var a abser
	f := myFloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f
	a = &v

	// error!
	//a = v

	fmt.Println(a.Abs())
}

type myFloat float64

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type vertex struct {
	X, Y float64
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
