import (
	"fmt"
)

// atribua uma função a uma variavel e a chame

func main() {
	y := func(x int) int {
		return (x * x * x)
	}

	fmt.Println(y(5))
}

