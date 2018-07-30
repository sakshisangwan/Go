package main 

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs" : {1.234, -34.567},
	"Google": {45.678, -56.78},
}

func main() {
	fmt.Println(m)
}