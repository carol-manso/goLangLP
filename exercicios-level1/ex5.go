import (
	"fmt"
)

type numeros int

var x numeros
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("Tipo x= %T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Printf("Valor y: %v, tipo y:%T", y, y)

}