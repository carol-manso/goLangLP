
import (
	"fmt"
)

// crie uma função que retorne uma função, atribuindo a func retornada a uma variavel. Chame a func retornada

func main() {
	y := retornaFuncao()
	fmt.Println(y(4))

}
func retornaFuncao() func(int) int {
	fmt.Println("Dentro da primeira função")
	return func(x int) int {
		return x * 4
	}

}
