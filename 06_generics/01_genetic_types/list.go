package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func addList[T any](l []*List[T], val T) []*List[T] {
	newItem := List[T]{val: val}
	l = append(l, &newItem)
	i := len(l)
	if i-1 > 0 {
		l[i-2].next = &newItem
	}
	return l
}

func main() {
	a := make([]*List[int], 0)

	for i := 1; i < 10; i++ {
		a = addList[int](a, i)
	}

	for l := a[0]; l != nil; l = l.next {
		fmt.Println(l.val)
	}
}
