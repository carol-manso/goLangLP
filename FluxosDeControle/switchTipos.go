import (
	"fmt"
)

var x interface{} //interface vazia
func main() {
	x = "ola"
	//switch vai ser em cima do tipo de x
	switch x.(type) {
	case int:
		fmt.Println("É um int")
	case bool:
		fmt.Println("É um bool")
	case string:
		fmt.Println("É uma string")
	case float64, float32:
		fmt.Println("É um float")
	}

}