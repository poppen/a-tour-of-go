package main

import (
	"fmt"
	"math"
)

func areApproximal(x, y float64) bool {
	if fmt.Sprintf("%.11f", x) == fmt.Sprintf("%.11f", y) {
		return true
	}
	return false
}

func Sqrt(x float64) (float64, int) {
	z := 1.0
	var p float64

	i := 1
	for ; i <= 10; {
		z -= (z * z - x) / (2 * z)
		if z == p || areApproximal(z, math.Sqrt(x)) {
			return z, i
		}
		p = z
		i++
	}
	return z, i
}

func main() {
	r, i := Sqrt(2)
	fmt.Printf("The answer is %v, it took %v times to calculate this.\n", r, i)
}
