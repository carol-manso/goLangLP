import (
	"fmt"
)

func main() {
	x := 10
	y := func(x int) int {

		return (x * 56)
	}
	fmt.Println(x, "vezes 56 é: ", y(x))

}
