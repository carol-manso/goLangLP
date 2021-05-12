import (
	"fmt"
)

const x = 10 //tipo indefinido até ser usado. Pode ser tanto int quanto float.
var y = 10   //no momento de compilação já define que é um int.
var c int
var d float64

const (
	z = 90
	w = 120
	h = 300
) //declarar várias constantes ao mesmo tempo

func main() {

	c = x //posso colocar o c=x ou então o k=x. O tipo de x é definido na hora de usar.
	d = x
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(z, w, h)
}
