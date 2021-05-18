
import (
	"fmt"
)

//defina constantes tipadas e não tipadas
const (
	x = 10
	y = 20
	z = 40
)
const w int = 50
const p float64 = 10.5

func main() {
	var k float64 = x
	fmt.Println(x, y, z, w, p, k)

}