package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	k, z := 0., 1.
	for {
		z, k = z - (z*z - x) / (2*z), z
		if math.Abs(k-z) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	i := 89. 
	fmt.Println(Sqrt(i))
	fmt.Println(Sqrt(i) == math.Sqrt(i))
}
