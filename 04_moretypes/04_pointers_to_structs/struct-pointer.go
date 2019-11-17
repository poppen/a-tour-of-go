package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func main() {
	v := vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
