import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10) //valor 0 como padrão
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10) //acima da minha capacidade.
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 11) //aumenta minha capacidade
	fmt.Println(slice, len(slice), cap(slice))
}
