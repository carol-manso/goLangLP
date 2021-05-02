
import (
	"fmt"
)

type hotdog int

var b hotdog = 10

func main() {
	fmt.Printf("Tipo da variável b: %T valor: %v", b, b)
	x := 10
	fmt.Printf("\n%v", x)
	//b=x -> erro de incompatibilidade de tipos.
	x = int(b)
	//agora consegui transformar o b em int. Conversão de tipos. Não coersão. Só temos conversão em Go.

}
