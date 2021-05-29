
import (
	"fmt"
)

func main() {
	ss := [][]int{
		//col:  0 1 2    //line
		[]int{1, 2, 3}, //0 //linha da matriz.Começa em 0
		[]int{4, 5, 6}, //1
		[]int{7, 8, 9}, //2
	}
	fmt.Println(ss)
	fmt.Println(ss[0][2])
	fmt.Println(ss[0])
	fmt.Println(ss[1])
	fmt.Println(ss[2])

}
