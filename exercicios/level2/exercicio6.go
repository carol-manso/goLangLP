

import (
	"fmt"
)

//usando iota, crie 4 constantes que indiquem os 4 próximos anos
const (
	y = iota + 2022
	z
	w
	l
	m
	n
)

func main() {

	fmt.Println(y, z, w, l, m, n)

}
