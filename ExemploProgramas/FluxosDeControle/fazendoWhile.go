//Em Go, não existe while, só existe o for. Porém, podemos transformar um for em while apenas retirando a condição e a pós ação do for

import (
	"fmt"
)

func main() {
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	//loop infinito:
	y := 0
	for {
		if y < 10 {
			y++
			fmt.Println("oi")
		} else {
			fmt.Println("tchau")
			break //apesar de ser um loop infinito, teoricamente, o break quebra-o, saindo do loop
		}
	}
	fmt.Println("fizemos um break no loop")
}
