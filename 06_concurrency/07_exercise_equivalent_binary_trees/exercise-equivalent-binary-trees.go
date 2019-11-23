package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		walk(t.Left, ch)
	}
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	v1 := make([]int, 10)
	v2 := make([]int, 10)

	go func() {
		walk(t1, c1)
		close(c1)
	}()
	go func() {
		walk(t2, c2)
		close(c2)
	}()

	for v := range c1 {
		v1 = append(v1, v)
	}
	for v := range c2 {
		v2 = append(v2, v)
	}

	if len(v1) != len(v2) {
		return false
	}

	sort.Sort(sort.IntSlice(v1))
	sort.Sort(sort.IntSlice(v2))

	for i := 0; i < len(v1); i++ {
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

	fmt.Printf("tree.New(1) and tree.New(1) are same?: %v\n", same(tree.New(1), tree.New(1)))
	fmt.Printf("tree.New(1) and tree.New(2) are same?: %v\n", same(tree.New(1), tree.New(2)))
}
