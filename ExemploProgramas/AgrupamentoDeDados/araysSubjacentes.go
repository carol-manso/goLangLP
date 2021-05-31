import (
	"fmt"
)

func main() {
	primeiroSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(primeiroSlice)
	segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...) //fatia do primeiro slice
	fmt.Println(segundoSlice)                                       //pegamos 1, 2 e 5
	fmt.Println(primeiroSlice)                                      //tem um 5 no lugar do 3!!

	//if it has sufficient capacity, the destination is resliced to accomodate the new elements;
	//if it doesn't fit, a new underlying array will be allocated
	//solução: copiar item por item ou pega o mesmo slice e vai modificando. Faz tudo na mesma variáveis ou copio item por item usando forloop.

}
