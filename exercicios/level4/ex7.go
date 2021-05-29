
import (
	"fmt"
)

//criar uma slice de 2 dimensões contendo nome, sobrenome e hobbie favorito; 3 pessoas incluidas; range para demosntrar os dados
func main() {
	pessoas := [][]string{
		[]string{"Cláudia", "Macedo", "Tocar violão"},
		[]string{"Sabrina", "Toledo", "Correr"},
		[]string{"Patrícia", "Klaus", "Jogar jogos de tabuleiro"},
	}
	fmt.Println(pessoas)
	for _, valor := range pessoas {
		for _, valorInterno := range valor {
			fmt.Print(valorInterno, "\t")
		}
		fmt.Println("\n===================")
	}

}
