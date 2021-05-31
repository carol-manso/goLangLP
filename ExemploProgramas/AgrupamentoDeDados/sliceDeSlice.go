
import (
	"fmt"
)

func main() {
	//0.           1.            2.                3.        4.
	sabores := []string{"calabresa", "margherita", "quatro queijos", "a moda", "espinafre"}
	//quero pegar apenas os 2 primeiros. Vou do 0 até o 2 , mas sem incluir o 2
	//0.       1.      2.      3.     4.
	fatia := sabores[0:2]
	for indice, valor := range fatia {
		fmt.Println(indice, valor)

	}
	fmt.Println("========")
	//do 2º ate o final.
	fatia = sabores[2:len(sabores)] //ou sabores [2:]
	for indice, valor := range fatia {
		fmt.Println(indice, valor)

	}

	fmt.Println("========")
	//do 0 até o 3
	fatia = sabores[:4]
	for indice, valor := range fatia {
		fmt.Println(indice, valor)

	}
	fmt.Println("========")

	fmt.Println(sabores)
	for i := 0; i < len(sabores); i++ {
		fmt.Println("*", sabores[i])
	}
}
