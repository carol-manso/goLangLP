import (
	"fmt"
)

//crie e utilize uma função anonima

func main() {
	x := 3
	func(x int) {
		fmt.Println(x * x * x)
	}(x)
}

