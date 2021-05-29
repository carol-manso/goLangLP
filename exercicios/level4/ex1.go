package main

import (
	"fmt"
)

//fazer uma literal composta para criar um array que suporte 5 valores do tipo int, atribuir valors aos índices e utilizar o range para mostrar os valores
func main() {
	array := [5]int{10, 200, 35, 45, 5}
	fmt.Println(array)
	for i, v := range array {

		fmt.Println(i, "->", v)

	}

	fmt.Printf("%T", array)

}
