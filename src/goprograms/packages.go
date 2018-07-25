// Programs start running in package main
// Import paths used by packages in this program: "fmt" and "math/rand"
// By convention, package name is same as last element in import path. 
// "math/rand" package comprises files beginning with statement "package rand"

package main

import ( 
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
}