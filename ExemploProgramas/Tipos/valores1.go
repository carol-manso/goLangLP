package main

import (
	"fmt"
)

var array = []int{0,1,2,3,4,5,6,7,8,9}
var slice  = array[5:]
var mapa = make(map[int]string)

//Da para rodar codigo, o erro é so por ter duas main no mesmo pacote
func main(){

	

	mapa[1] ="Teste"
	mapa[4] = "Correto"

	fmt.Println(array)
	fmt.Println(slice,"\n")

	slice2 := make([]string,4)
	
	slice2 = []string{"Joao","Carol","Marco","Claudio"}

	fmt.Println(slice2)

	fmt.Println("\n",mapa[4])
	fmt.Println("",mapa[1])

	slice2 = nil	//nao necessario pois Golang tem garbage collector
	


}
