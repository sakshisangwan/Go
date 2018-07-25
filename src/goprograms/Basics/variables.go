// type is written at last of the var statement
// var statement can be at package or function level


package main

import "fmt"

var c, python, java bool		//package level

func main() {
	var i int					// function level
	fmt.Println(i, c, python, java)
}