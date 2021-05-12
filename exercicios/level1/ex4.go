

import (
	"fmt"
)

type numeros int

var x numeros

func main() {
	fmt.Println(x)
	fmt.Printf("Tipo x= %T\n", x)
	x = 42
	fmt.Println(x)

}
