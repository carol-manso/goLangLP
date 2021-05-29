

import (
	"fmt"
)

// programa para demostrar a declaração if(), else if() e else
func main() {
	n := 420
	if n <= 200 {
		fmt.Println("Sou menor que 201")
	} else if n <= 300 {
		fmt.Println("Sou menor que 301")
	} else {
		fmt.Println("Sou maior que 300")
	}
}
