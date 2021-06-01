import (
	"fmt"
)

// demonstre o funcionamento de um callback (função que recebe outra função

func main() {
	recebeFuncao(argumento)

}

func argumento() {
	fmt.Println("Sou um argumento!")
}
func recebeFuncao(f func()) {
	fmt.Println("Estou recebendo uma função como argumento")
	f()

}