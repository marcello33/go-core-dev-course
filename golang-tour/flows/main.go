package main

import (
	"fmt"
	"math"
)

// Sqrt implements the square root function using the Newton's technique
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z = z - (z*z-x)/(2*z)
	}

	return z
}

func main() {
	v := 33.0
	fmt.Printf("Using Newton technique: %f\nUsing build-in method: %f", Sqrt(v), math.Sqrt(v))
}
