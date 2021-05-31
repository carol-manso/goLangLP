

import (
	"fmt"
)

//- Criar um tipo "pessoa" com tipo subjacente struct, que possa conter nome e sobrenome e sabor favorito de sorvete como campo
//demostrar sorvetes  com range (slice of string)
type pessoa struct {
	nome       string
	sobrenome  string
	sorveteFav []string
}

func main() {
	pessoa1 := pessoa{"Lincon", "Fernandes", []string{"chocolate", "limão", "carambola"}}
	pessoa2 := pessoa{nome: "Gabriel", sobrenome: "Silva", sorveteFav: []string{"morango", "banana", "chocolate"}}
	fmt.Println(pessoa1.nome, pessoa1.sobrenome, "\n Sorvetes preferidos:")
	for _, valor := range pessoa1.sorveteFav {
		fmt.Println(valor)

	}
	fmt.Println("===============")
	fmt.Println(pessoa2.nome, pessoa2.sobrenome, "\n Sorvetes preferidos:")
	for _, valor := range pessoa2.sorveteFav {
		fmt.Println(valor)

	}

}
