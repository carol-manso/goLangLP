import (
	"fmt"
)

// demonstre o funcionamento de um closure - função que recebe outra função, onde ela tem uma var alem de seu escopo

func main() {
	y := retornaFuncao()
	y()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	z := retornaFuncao()
	fmt.Println(z())
	fmt.Println(z())

}

func retornaFuncao() func() int {
	fmt.Println("Estou recebendo uma função como argumento")
	x := 10
	return func() int {
		x += 10
		return x
	}

}