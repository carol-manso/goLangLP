
import (
	"fmt"
)

var x int = 42
var y string = "James B."
var z bool = true

func main() {
	s := fmt.Sprint(x, " ", y, " ", z) //junto x e y em uma string, formata uma string
	fmt.Println(s)
}
