package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x) //zero value
	x = true
	fmt.Println(x)      //atribuindo valores
	fmt.Println(1 == 1) //op relacionais
	fmt.Println(10 > 1)
	fmt.Println(10 >= 100)
}
