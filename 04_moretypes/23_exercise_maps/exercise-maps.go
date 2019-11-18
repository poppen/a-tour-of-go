package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func wordCount(s string) map[string]int {
	r := make(map[string]int)

	for _, v := range strings.Fields(s) {
		r[v]++
	}

	return r
}

func main() {
	wc.Test(wordCount)
}
