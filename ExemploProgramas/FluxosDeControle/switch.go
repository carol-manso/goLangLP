
import (
	"fmt"
)

func main() {
	pessoaNoEscritorio := "Pedro"
	x := 8
	if (x > 5) == true {
		fmt.Println("x é maior que 5")
	}

	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case (x == 4), (x == 8):
		fmt.Println(" x é 4 ou 8")
	case x > 5:
		fmt.Println("x é maior que 5")
	}
	switch pessoaNoEscritorio {
	case "Marcos":
		fmt.Println("Marcos está no escritório")
	case "Ju", "Pedro", "Carol":
		fmt.Println("time 1 está na firma")

	case "Joana":
		fmt.Println("Joana está no escritório")
		fallthrough
	case "Maria":
		fmt.Println("Maria vem sempre que Joana vem")

	default:
		fmt.Println("Está vazio, ou temos um feriado ou todos estão de ressca ")
	}
}
