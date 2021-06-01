import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) saudacao() {
	fmt.Println("oi,", p.nome, ",bom dia!")
}

func main() {
	Mauricio := pessoa{"Maurício de Souza", 35}
	Paula := pessoa{"Paula de Souza", 35}
	Mauricio.saudacao()
	Paula.saudacao()
}
