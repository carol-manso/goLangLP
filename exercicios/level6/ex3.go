import (
	"fmt"
)

//utilize defer para demostrar sua exec apenas depois do contexto da qual é pertencente
func main() {
	slice := []int{1, 2, 3, 4, 5}
	defer fmt.Println(slice)
	fmt.Println("oi")
}

