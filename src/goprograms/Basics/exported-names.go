// A name is exported if it beigns with a capital letter.
// Here Pi is exported from math package (using pi would result in error)


package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
