
import (
	"fmt"
)

func main() {
	k := somenteImpares(soma, []int{1, 2, 5, 89, 63, 54, 52, 344, 22, 28, 30}...)
	fmt.Println(k)
}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
func somenteImpares(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, value := range y {
		if value%2 != 0 {
			slice = append(slice, value)
		}
	}
	total := f(slice...)
	return total
}