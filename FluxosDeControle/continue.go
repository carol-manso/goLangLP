import (
	"fmt"
)

func main() {
	x := 0
	for x <= 20 {
		x++
		if x%2 != 0 {
			//para aqui e vai para a próxima iteração se for ímpar.
			continue
		}
		fmt.Pritln(x)
	}
}

//break pega o loop inteiro e joga fora. Já o continue quebra a iteração atual. Se quiser parar no meio de uma interação e ir para a próxima, uso o continue. 