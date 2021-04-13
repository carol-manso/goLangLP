package main

import (
	"fmt"
)

var x int
var y float64
var z string
var w bool

func main() {
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", z, z)
	fmt.Printf("%v, %T\n", w, w)
}
