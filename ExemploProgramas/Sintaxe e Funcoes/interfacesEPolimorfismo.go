import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}
type dentista struct {
	pessoa
	dentesArrancados int
	salário          float64
}
type arquiteto struct {
	pessoa
	tipoDeConstrucao string
	tamanhoDaLoucura int
}

func (x dentista) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, ",já arranquei", x.dentesArrancados, "dentes. Bom dia!")
}
func (x arquiteto) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, ".Bora fazer um projeto de ", x.tipoDeConstrucao)
}

type gente interface {
	oiBomDia()
} //oiBomDia é o método que preciso para implementar aquela interface.

func serHumano(g gente) {
	g.oiBomDia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho", g.(dentista).salário)
	case arquiteto:
		fmt.Println("Eu construo", g.(arquiteto).tipoDeConstrucao)

	}

}
func main() {
	Robson := arquiteto{pessoa{"Roberto", "Gentleman", 35}, "predio", 5}
	Leo := dentista{pessoa{"Adriano", "Estevo", 42}, 15, 2500}
	//Robson.oiBomDia()
	//Leo.oiBomDia()
	serHumano(Robson)
	serHumano(Leo)
	//funciona igual. Ao invés de usar um método específico de cada tipo, consigo ter uma interface para tipos diferentes que chama a mesma coisa dentro de cada valor.

} 