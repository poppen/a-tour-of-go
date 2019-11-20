package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is float64
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var p float64
	z := 1.0
	for {
		z -= (z*z - x) / (2 * z)

		if p == z || math.Abs(p-z) < 0.1e-11 {
			return z, nil
		}

		p = z
	}
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
