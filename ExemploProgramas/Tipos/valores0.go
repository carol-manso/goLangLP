package main

import (
	"fmt"
)
var v *int
var w bool
var x int
var y float64
var z string

func main() {
	v = new(int)
	*v = 4

	fmt.Printf("Valor = %v\t Tipo = %T\t Endereço = %v\n", *v,v,v)
	fmt.Printf("Valor = %v\t Tipo = %T\n", w, w)
	fmt.Printf("Valor = %v\t Tipo = %T\n", x, x)
	fmt.Printf("Valor = %v\t Tipo = %T\n", y, y)
	fmt.Printf("Valor = %v\t Tipo = %T\n", z, z)
	
}
