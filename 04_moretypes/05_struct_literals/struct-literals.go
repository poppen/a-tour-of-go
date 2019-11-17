package main

import "fmt"

type vertex struct {
	X, Y int
}

var (
	v1 = vertex{1, 2}
	v2 = vertex{X: 1}
	v3 = vertex{}
	p  = &vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
