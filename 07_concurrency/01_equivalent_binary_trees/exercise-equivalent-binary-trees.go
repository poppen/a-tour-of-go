package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, depth int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch, depth+1)
	}
	if t.Right != nil {
		Walk(t.Right, ch, depth+1)
	}

	if depth == 0 {
		close(ch)
	}
}

func Sort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	high := make([]int, 0)
	low := make([]int, 0)

	k := a[len(a)-1]
	for i := 0; i < len(a)-1; i++ {
		if a[i] <= k {
			low = append(low, a[i])
		} else {
			high = append(high, a[i])
		}
	}
	return append(append(Sort(low), k), Sort(high)...)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	a1 := make([]int, 0)
	a2 := make([]int, 0)

	go Walk(t1, ch1, 0)
	go Walk(t2, ch2, 0)

	for v1 := range ch1 {
		a1 = append(a1, v1)
	}
	for v2 := range ch2 {
		a2 = append(a2, v2)
	}

	a1 = Sort(a1)
	a2 = Sort(a2)
	if len(a1) != len(a2) {
		return false
	}
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch, 0)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
