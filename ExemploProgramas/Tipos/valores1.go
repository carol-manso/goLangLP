package main

import (
	"fmt"
)

//Variáveis globais
var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

var slice = array[5:]

var mapa = make(map[int]string)

func main() {

	mapa[1] = "Teste"
	mapa[4] = "Correto"

	fmt.Println(array)
	fmt.Println(slice)

	//Inferência de tipo
	slice2 := make([]string, 4)
	slice2 = []string{"Joao", "Carol", "Marco", "Claudio"}

	fmt.Println(slice2)

	fmt.Println("\n", mapa[4])
	fmt.Println("", mapa[1])

	slice2 = nil

}
