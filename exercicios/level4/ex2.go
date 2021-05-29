import (
	"fmt"
)

//fazer um slice do tipo int e atribuir 10 valores a ela, usar range e demostrar seu tipo
func main() {
	slice := []int{10, 200, 35, 45, 5, 78, 98, 65, 3, 86}
	fmt.Println(slice)
	for i, v := range slice {

		fmt.Println(i, "->", v)

	}

	fmt.Printf("%T", slice)

}
