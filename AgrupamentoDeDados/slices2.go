import (
	"fmt"
)

func main() {
	slice := []string {"banana", "laranja", "maçã", "jaca", "pessego" }
	for indice, valor := range slice{
		fmt.Println("No índice", indice, " temos o valor:", valor)
	}
	
	slice = append(slice,"melancia")
	
	for _, valor := range slice{
		fmt.Println("Valores:", valor)
	}
	
	sliceInt:= []int {25,30,35,40,68,95,36,70}
	total:=0
	for _, valor := range sliceInt{
		total+=valor
	}
	fmt.Println("soma dos valores= ", total )
}