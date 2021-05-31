import (
	"fmt"
)

type pessoa struct {
	nome       string
	sobrenome  string
	sorveteFav []string
}

// Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// Demonstre os valores do map utilizando range.
func main() {
	pessoa1 := pessoa{"Lincon", "Fernandes", []string{"chocolate", "limão", "carambola"}}
	pessoa2 := pessoa{nome: "Gabriel", sobrenome: "Silva", sorveteFav: []string{"morango", "banana", "chocolate"}}

	mapPeople := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for _, valor := range mapPeople {
		fmt.Println(valor.nome, valor.sobrenome)
		for _, value := range valor.sorveteFav {
			fmt.Println(value)
		}
		fmt.Println("============")

	}

}
