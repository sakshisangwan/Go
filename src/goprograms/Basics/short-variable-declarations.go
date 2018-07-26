// := is the short assignment
// Outside a function the short assignment cannot be used


package main

import "fmt"

func main() {
	var i int = 12
	j := 3
	k, l, m := true, false, "no!"

	fmt.Println(i, j, k, l, m)
}