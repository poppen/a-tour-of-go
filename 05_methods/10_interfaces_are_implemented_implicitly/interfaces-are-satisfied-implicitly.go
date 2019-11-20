package main

import (
	"fmt"
)

// I is interface
type I interface {
	M()
}

// T is struct
type T struct {
	S string
}

// M prints T.S
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
