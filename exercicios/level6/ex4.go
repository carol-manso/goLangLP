import (
	"fmt"
)

//struct pessoa com nome, sobrenome, e idade
//método para demostrar nome e idade
//demostre esse método e o seu resultado

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) MostrarPessoa() {
	fmt.Println("Oi, meu nome é", p.nome, "e tenho", p.idade, " anos de idade")
}
func main() {
	Marcos := pessoa{"Marcos Aurélio de Souza", 65}
	Marcos.MostrarPessoa()
}
