

import (
	"fmt"
)

// programa para demostrar switch sem statement especificado- qualquer coisa que for true está valendo

func main() {
	n := 420
	switch {
	case n <= 100:
		fmt.Println("Sou um numero de 0 a 100")
	case n > 100 && n <= 200:
		fmt.Println("Sou um numero de 100 a 200")
	case n > 200:
		fmt.Println("Sou um numero maior que 200")
	}
}
