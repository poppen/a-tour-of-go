package main

import "fmt"

func fibonacci() func() int {
	beforeTwo := 1
	prev := 0
	cur := 0

	return func() int {
		cur = beforeTwo + prev
		beforeTwo, prev = prev, cur
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
