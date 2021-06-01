import (
	"fmt"
)

func main() {
	x := 10
	func(x int) {
		fmt.Println(x, "vezes 56 é: ")
		fmt.Println(x * 56)
	}(x)
}
