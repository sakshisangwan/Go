// Return values can be named, they are treated as variables
// A return statement without arguments returns named return values.


package main 

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}