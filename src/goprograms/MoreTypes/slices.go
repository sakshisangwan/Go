// Inludes the low bound and excludes the high bound
// a[1 : 4] elements 1 through 3 of a are sliced


package main 

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}