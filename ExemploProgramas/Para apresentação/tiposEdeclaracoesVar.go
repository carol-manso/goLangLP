package main

import (
	"fmt"
)

var IntVar int
var FloatVar float32 //= 65.5
var StringVar string
var BoolVar bool //= false

func main() {
	fmt.Println(IntVar)
	fmt.Println(FloatVar)
	fmt.Println(StringVar)
	fmt.Println(BoolVar)
	IntVar = 65
	StringVar = "Olá, pessoal! Vamos aprender Go"
	fmt.Printf("valor de IntVar: %v\ntipo de IntVar: %T\nIntVar na tabela ASCII: %#U \nIntVar na base 2 : %b \n", IntVar, IntVar, IntVar, IntVar)
	fmt.Println("=======================")
	fmt.Printf("valor de FloatVar: %v\ntipo de FloatVar: %T\n", FloatVar, FloatVar)
	fmt.Println("=======================")
	fmt.Printf("valor de StringVar: %v\ntipo de StringVar: %T\n", StringVar, StringVar)
	fmt.Println("=======================")
	fmt.Printf("valor de BoolVar: %v\n tipo de BoolVar: %T\n", BoolVar, BoolVar)
	fmt.Println("=======================")
	var x = 10
	var y int = 10
	fmt.Println(x)
	fmt.Println("=======================")
	fmt.Println(y)

}
