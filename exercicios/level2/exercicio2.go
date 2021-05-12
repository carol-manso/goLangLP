import (
	"fmt"
)

//fazer expressões com operadores ==, != , <=, < , >, >=

func main() {

	x := 3
	y := 5
	z := true
	w := x == y
	l := x <= y
	m := x < y
	x = 70
	y = 23
	n := y >= x
	o := y > x
	p := y != x
	fmt.Println(x, y, z, w, l, m, n, o, p)

}