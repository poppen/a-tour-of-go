package main

import "fmt"

// I is interface
type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // causes runtime error
}

func describe(i I) {
	fmt.Printf("%v, %T\n", i, i)
}
