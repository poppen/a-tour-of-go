package main

import (
	"fmt"
)

func main() {
	q := []int{2, 3, 5, 7, 1, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, false},
	}
	fmt.Println(s)
}
