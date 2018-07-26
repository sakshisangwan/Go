// If not specified the type of variable is inferred fromvalue on right


package main

import "fmt"

func main () {
	a := 112
	b := 23.45
	c := false

	fmt.Printf("The type of a is %T\n", a)
	fmt.Printf("The type of b is %T\n", b)
	fmt.Printf("The type of c is %T\n", c)
}