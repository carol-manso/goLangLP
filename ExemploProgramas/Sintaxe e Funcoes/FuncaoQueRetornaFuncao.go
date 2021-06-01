import (
	"fmt"
)

func main() {
	x := retornaFuncao() //função que retorna uma função. Esse retorno se torna o valor de x
	y := x(30)           //passo o arguento 30 para a função.
	fmt.Println(y)
	fmt.Println(retornaFuncao()(3)) //forma direta de fazer
}
func retornaFuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
