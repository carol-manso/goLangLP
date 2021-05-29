
import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice2 := append(slice, 6)
	fmt.Println(slice2)
	fmt.Println(slice[3])
	slice[3] = 345615
	fmt.Println(slice[3])
	//slice[20]=30 //-> erro! Não pode colocar em posição que não segue a ordem ou sem append
}
