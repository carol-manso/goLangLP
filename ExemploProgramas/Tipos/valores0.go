package main

import (
	"fmt"
)
var u rune 		//mesmo que char, mas usando unicode. Alias de int32
var v *int
var w bool
var x int		//pode ser trocado por int8, int16,int32,int64
var y float64	//pode ser trocado por float32
var z string


//Da para rodar codigo, o erro é so por ter duas main no mesmo pacote
func main() {
	v = new(int)
	*v = 4 
	x = 3
	fmt.Printf("Valor = %c\t Tipo = %T\n", u, u)
	fmt.Printf("Valor = %v\t Tipo = %T\t Endereço = %p\n", *v,v,v)
	fmt.Printf("Valor = %v\t Tipo = %T\n", w, w)
	fmt.Printf("Valor = %v\t Tipo = %T\n", x, x)
	fmt.Printf("Valor = %v\t Tipo = %T\n", y, y)
	fmt.Printf("Valor = %v\t Tipo = %T\n", z, z)


	
}
