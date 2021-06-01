package main

import (
	"fmt"
)

func main() {
	basica()
	argumento("tarde")
	valor := somaDeDois(3, 5)
	fmt.Println(valor)
	valor2, quantidade := soma(3, 5, 7, 9) //duas variáveis para 2 retornos.
	fmt.Println(valor2, quantidade)
	saudacao, media := mediaSaudacao("manhã", 15, 8, 5)
	fmt.Println(saudacao, media)

}

func basica() {
	fmt.Println("Estou na função básica!")
}

func argumento(x string) {
	if x == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if x == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}

} //função com  argumentos

//x int, y int OU x,y int -> mesmo tipo
func somaDeDois(x, y int) int {
	return x + y
} //função com retorno

func soma(x ...int) (int, int) {
	sum := 0
	for _, valor := range x {
		sum += valor
	}
	return sum, len(x)
} //valor variadico de ints(recebe uma quantidade indefinida de ints) + varios retornos -> retorna a soma e a quantidade de elementos recebidos.

func mediaSaudacao(s string, x ...int) (string, float64) {
	saudacao := ""
	if s == "manhã" {
		saudacao = "Oi, bom dia!"
	} else if s == "tarde" {
		saudacao = "Oi, boa tarde!"
	} else {
		saudacao = "Oi, boa noite!"
	}
	saudacao += " O resultado da sua média é:"
	soma := 0
	for _, value := range x {
		soma += value
	}
	media := float64(soma) / float64(len(x))
	return saudacao, media
}

//parâmetro variádico SEMPRE deve ser o último.
