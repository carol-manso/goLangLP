package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	for i := 0; i < 5; i++ {
		x[i] = i
		fmt.Println(x[i])
	}
	fmt.Println(x) //posso printar o array inteiro de uma vez
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y) //TIPOS DE X E Y INCOMPATÍVEIS
	fmt.Println(len(y))   //tamanho
}
