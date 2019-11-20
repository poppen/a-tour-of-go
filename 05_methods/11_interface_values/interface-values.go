package main

import (
	"fmt"
	"math"
)

// I is interface
type I interface {
	M()
}

// T is struct
type T struct {
	S string
}

// M is function
func (t *T) M() {
	fmt.Println(t.S)
}

// F is float64
type F float64

// M is function
func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
