package main

import (
	"fmt"
)

// x:=10 -> ERRO!
func main() {

	intInferencia := 10
	floatInferencia := 25.6
	stringInferencia := "String"
	boolInferencia := false
	fmt.Printf("valor de intInferencia: %v\ntipo de intInferencia: %T\n intInferencia na tabela ASCII: %#U\n intInferencia na base 2 : %b \n", intInferencia, intInferencia, intInferencia, intInferencia)
	fmt.Println("=======================")
	fmt.Printf("valor de floatInferencia: %v\n tipo de floatInferencia: %T\n", floatInferencia, floatInferencia)
	fmt.Println("=======================")
	fmt.Printf("valor de stringInferencia: %v\n tipo de stringInferencia: %T\n", stringInferencia, stringInferencia)
	fmt.Println("=======================")
	fmt.Printf("valor de boolInferencia: %v\ntipo de boolInferencia: %T\n", boolInferencia, boolInferencia)
	fmt.Println("=======================")

}
