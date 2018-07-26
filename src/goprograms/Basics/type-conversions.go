// Assignment between different types requires explicit conversion


package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 1, 2
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}