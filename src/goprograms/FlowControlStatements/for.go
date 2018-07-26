// Go has only one looping construct, the for loop.

// The basic for loop has three components separated by semicolons:

// the init statement: executed before the first iteration
// the condition expression: evaluated before every iteration
// the post statement: executed at the end of every iteration


package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// The init and post statements are optional

	sum1 := 1
	for ; sum1 <1000; {
		sum1 +=sum1
	}
	fmt.Println(sum1)
}