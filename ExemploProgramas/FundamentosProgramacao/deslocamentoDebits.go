import (
	"fmt"
)

const (
	_  = iota             //0
	KB = 1 << (iota * 10) //1 << (1*10) -> 1 em binário deslocado 10 casas para a esquerda
	MB                    //1 << (2*10) -> 1 em binário deslocado 20 casas para a esquerda
	GB                    //1 << (3*10) -> 1 em binário deslocado 30 casas para a esq
	TB                    //1 << (4*10)
)

func main() {
	y := 24
	x := y << 2
	z := y >> 2 //desloco 2 bits para a direita.
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)
	fmt.Println("================================")
	fmt.Println("Conversão de bytes")
	fmt.Printf("binary \t\t\t\t decimal\n")
	fmt.Printf("%b \t\t\t", KB)
	fmt.Printf("%d \n", KB)
	fmt.Printf("%b \t\t", MB)
	fmt.Printf("%d \n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b", TB)
	fmt.Printf("  %d \t  \n", TB)

}
