package main

import (
	"fmt"
)

func main() {
	x := "oi, bom dia\n como vai? \n \" que \" " ///string interpretada-> \n, \" São as interpret strings literalsão interpretados.
	y := `oiii              \n      \" blz?`     // raw = cru. Não interpreta nada do que está lá dentro.

	fmt.Println(x)
	fmt.Println(y)
	x = "oi"
	y = "bom dia"
	fmt.Print(x)               //não tem uma linha nova
	z := fmt.Sprint(x, " ", y) //junto x e y em uma string
	fmt.Println(z)
}
