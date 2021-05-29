import (
	"fmt"
)

// resto da div dos numeros entre 10 e 100 por 4
func main() {
	n := 4
	for i := 10; i <= 100; i++ {
		fmt.Println(i % n)

	}
}