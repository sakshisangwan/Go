// & operator generates a pointer to its operand
// * operator denotes the pointer's underlying value
// dereferencing means directly accessing the value through a pointer 


package main 

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p =21				//dereferencing
	fmt.Println(i)

	p = &j
	*p = *p/37			//dereferencing
	fmt.Println(j)
}