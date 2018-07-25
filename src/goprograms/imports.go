// Grouping imports into a parenthesized, "factored" import statement
// Can be written like:
// import "fmt"
// import "math"
// But factored import statement is a better approach


package main

import (
	 "fmt"
	 "math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(16))
}