
import (
	"fmt"
)

//Usar o slice do ex anterior e fati=a-lo -> 1º ao 3º;1º ao 5º; 2º ao 7º , 3º ao penultimo; 3º ao ultimo usando len()
func main() {
	//            0   1  2  3  4  5  6   7  8  9
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2:(len(slice) - 1)])
}
