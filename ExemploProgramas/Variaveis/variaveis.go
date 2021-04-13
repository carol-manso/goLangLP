package main

import (
	"fmt"
)

var y = "Olá, bom dia!"

func main() {
	x := 10

	fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	x, z := 20, 30 //só se coloca := para atribuição de novo valor se tiver uma nova variável sendo declarada.
	fmt.Printf("x: %v, %T, z: %v, %T \n", x, x, z, z)
}
