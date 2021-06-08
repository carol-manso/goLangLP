package main

import (
	"fmt"
)

//utilize defer para demostrar sua exec apenas depois do contexto da qual é pertencente
func main() {
	nomes := []string{"Ana Carolina", "Cláudio", "João Victor", "Marco Tulio"}
	defer fmt.Println(nomes)
	fmt.Println("Nosso grupo é composto por: ")
}
