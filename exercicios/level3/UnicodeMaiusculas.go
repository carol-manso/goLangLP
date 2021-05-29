

import (
	"fmt"
)

//printar na tela o unicode point de todas as letras maiúsculas, 3x cada
func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v \n", i)
		for k := 0; k < 3; k++ {
			fmt.Printf("%#U \n", i) //unicode code point
		}
	}
}
