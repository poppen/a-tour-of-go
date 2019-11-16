package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func main() {
	v := vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
