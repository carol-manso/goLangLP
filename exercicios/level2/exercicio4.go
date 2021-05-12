import (
	"fmt"
)

//demonstrar deslocamento de 1 bit para a esquerda

func main() {

	var x int = 32
	fmt.Printf("%d, %b, %#x", x, x, x)
	y := x << 1
	fmt.Printf("\n%d, %b, %#x", y, y, y)

}
