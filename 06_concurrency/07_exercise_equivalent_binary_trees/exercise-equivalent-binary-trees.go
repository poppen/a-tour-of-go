package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		go walk(t.Left, ch)
	}
	if t.Right != nil {
		go walk(t.Right, ch)
	}
}

func same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	v1 := make([]int, 10)
	v2 := make([]int, 10)

	go walk(t1, c1)
	go walk(t2, c2)

	for i := 0; i < 10; i++ {
		v1[i] = <-c1
		v2[i] = <-c2
	}

	if len(v1) != len(v2) {
		return false
	}

	sort.Sort(sort.IntSlice(v1))
	sort.Sort(sort.IntSlice(v2))

	for i := 0; i < 10; i++ {
		if v1[i] != v2[i] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)

	go walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("%v\n", same(tree.New(1), tree.New(1)))
}
