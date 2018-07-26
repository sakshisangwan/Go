// Constants are declared using "const" keyword (cannot use :=)
// Constants can be char, bool, stirng, numeric values

package main

import "fmt"

const Pi = 3.14

func main() {
	World := "there!"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	Truth := true
	fmt.Println("Go rules?", Truth)
}