// Function to add two numbers


// Function can take 0 to any number of arguments
// type comes after variable name like: "x int"


package main
 
import "fmt"

func add(x int, y int) int {		//Function1
	return x + y
}

// When 2 or more function parameters share a type, type can be omitted from all but the last parameter.

func add1(x, y, z int) int {		//Function2
	return x + y + z
}

func main() {
	fmt.Println(add(12,13))

	fmt.Println(add1(12,13,14))
}