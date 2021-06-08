import (
	"fmt"
)

func main() {
	canal := make(chan int, 1)
	go func() {
		canal <- 42 //coloco uma info no meu canal
	}()

	fmt.Println(<-canal) //goroutine querendo retirar info do canal.
} 