package main

import (
	"bufio"
	"fmt"
	"os"
)




func main() {
	
	fmtEscrita()
	//fmtLeitura()
	
	
}

func fmtEscrita() {

	x := "oi, bom dia\n como vai? \n \" que \" " ///string interpretada-> \n, \" São as interpret strings literal são interpretados.
	y := `oiii              \n      \" blz?`     // raw = cru. Não interpreta nada do que está lá dentro.

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("Olaa %v\n","Bom dia!")
	x = "oi"
	y = "bom dia"
	fmt.Print(x)               //não tem uma linha nova
	z := fmt.Sprint(x, " ", y) //junto x e y em uma string, formata uma string
	
	fmt.Println(z)


}


func fmtLeitura(){

	reader := bufio.NewReader(os.Stdin)

	
	fmt.Print("Digite um caractere: ") 
	entrada, _, _ := reader.ReadRune()	//le o caractere
	reader.Discard(1)					//descarta o '\n' ao fim da linha
	
	//A leitura pode ocorrer de dois metodos diferentes
	
	fmt.Print("Digite um numero: ")
	//input,_ := reader.ReadString('\n')	//Le a string incluindo o '\n'
	//input = strings.TrimSpace(input)		//remove espaços e quebras de linhas
	//inteiro,_ := strconv.Atoi(input)		//converte a string para inteiro

	var inteiro int
	fmt.Scanf("%d\n",&inteiro)

	fmt.Printf("entrada = %c inteiro = %d\n",entrada,inteiro)





}
