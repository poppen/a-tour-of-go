package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	i := 0

	var p float64

	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v %v\n", i, z)

		if p == z || math.Abs(p-z) < 0.1e-11 {
			return z, i
		}

		p = z
		i++
	}
}

func main() {
	z, i := sqrt(2)
	fmt.Printf("The answer of root square %v is %v. Tried %v times.", 2, z, i)
}
