import (
	"fmt"
)

var y = 25

func main() {
	fmt.Println("Hello!")
	y := 10
	qualquer(y)
}
func qualquer(x int) {
	fmt.Println(x)
	fmt.Println(y)
}