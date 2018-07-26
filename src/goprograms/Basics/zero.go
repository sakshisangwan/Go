// If not declared variables are explicitly initialized to 0
// numeric -> 0	 	boolean -> false		strings -> "" (empty string)

package main

import "fmt"

func main() {
	var i int 
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}