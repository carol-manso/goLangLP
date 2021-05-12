
import (
	"fmt"
)

const (
	a = iota //a=iota*5
	_
	b
	_ //descartar valores para eles continuarem a serem contados no iota, mas não computados.  EX: jogar 0 fora
	c
	x
	y
	z
) //declarar várias constantes ao mesmo tempo

func main() {
	fmt.Println(a, b, c, x, y, z) //faz sozinho para mim inteiros sucessivos não tipados.

}