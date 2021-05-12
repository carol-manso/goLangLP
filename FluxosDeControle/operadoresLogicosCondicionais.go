

import (
	"fmt"
)

func main() {
	x := 6
	if x % 2 == 0 || x % 5 == 0 {
		fmt.Println("x é multiplo de 5 ou de 2 ")
	} //basta uma ser true para dar true

	if x % 2 == 0 && x % 3 == 0 {
		fmt.Println("x é multiplo de 6")
	} //basta uma ser ser falsa para dar  false
}