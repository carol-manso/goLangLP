
import (
	"fmt"
)

func main() {
	sliceInts := []int{10, 20, 30, 40, 50}
	//fmt.Println( soma(sliceInts) ) -> ERRO- tipos diferentes, não posso converter um slice de ints para um int
	fmt.Println(soma(10, 20, 30, 40, 50))
	fmt.Println(soma(sliceInts...)) //operador de enumeração : ...
	fmt.Pritnln(soma())

}

func soma(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}


