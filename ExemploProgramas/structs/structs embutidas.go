import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}
type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {
	pessoa1 := pessoa{"Giovanni", 22}
	profissional1 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "pizzaiola",
		salario: 10000,
	}
	fmt.Println(pessoa1, profissional1)

	//selecao de valores:

	fmt.Println(pessoa1.nome, ",", pessoa1.idade, "\n====\n", profissional1.pessoa.nome, profissional1.titulo)
	profissional2 := profissional{pessoa{"Juliano", 35}, "vendedor", 2000}
	fmt.Println(profissional2)
	fmt.Println(profissional2.nome)
}
