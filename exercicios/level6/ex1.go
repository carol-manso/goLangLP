package main

import (
	"fmt"
)

//crie uma função que retorne um int; depois uma função que retorna um int e uma string; chamea-as e demonstre o result
func main() {
	fmt.Println(soma(1, 2, 3, 4, 5))
	texto, num := SaudacaoSoma(1, 2, 3, 4, 5)
	fmt.Println(texto, num)
}

func soma(x ...int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}
func SaudacaoSoma(x ...int) (string, int) {
	total := soma(x...)
	return "O total da soma desejada é:", total
}
