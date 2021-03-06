package main

import "fmt"

const (
	Big = 1 << 100 // Shifting a 1 bit by 100 places
	Small = Big >> 99 //Shifting it right again by 99 places i.e 1<<1, or 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))	-> const overflows int
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}