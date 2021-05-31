package main

import (
	"fmt"
)

/* tipo: veículo; campos: portas, cor + dois novos tipos: caminhonete e sedan
    - Ambos devem conter "veículo" como struct embutido
    caminhonete: campo bool chamado "traçãoNasQuatro"
    sedan deve: campo bool chamado "modeloLuxo"
    - Usando composite literal para caminhonete e sedan
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.*/

type veiculo struct {
	portas int
	cor    string
}
type caminhonete struct {
	veiculo
	tracao4Rodas bool
}
type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	caminhonetes := []caminhonete{
		caminhonete{
			veiculo:      veiculo{4, "prata"},
			tracao4Rodas: true,
		},
		caminhonete{veiculo{2, "vermelho"}, false},
	}
	sedans := []sedan{
		sedan{veiculo{4, "preto"}, true},
		sedan{veiculo{4, "azul"}, false},
	}
	fmt.Println("caminhonetes:")
	for indice, valor := range caminhonetes {
		fmt.Println("caminhonete ", indice+1, ":")
		fmt.Println(valor.portas, "portas e", valor.cor)
		fmt.Println("Tração nas 4 rodas:", valor.tracao4Rodas)
		fmt.Println("============")
	}

	fmt.Println("sedans:")
	for indice, valor := range sedans {
		fmt.Println("sedan ", indice+1, ":")
		fmt.Println(valor.portas, "portas e", valor.cor)
		fmt.Println("modelo luxo:", valor.modeloLuxo)
		fmt.Println("============")
	}

}
