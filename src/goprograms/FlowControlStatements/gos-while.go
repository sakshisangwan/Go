// for behaves as while on dropping the init, post statements as well as the semi-colons 

package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}