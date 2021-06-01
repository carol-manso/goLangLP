
import (
	"fmt"
)

//fazer uma struct anonima com um map e um slice.
func main() {
	defer fmt.Println("com defer, está antes de sem defer no código") //deixa pra última hora.
	fmt.Println("sem defer, está depois de com defer no código")
	defer fmt.Println(1)
	fmt.Println(2, 3, 4, 5)

}



