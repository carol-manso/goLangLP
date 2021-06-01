import (
	"fmt"
)

func main() {
	k := i()
	fmt.Println(k()) //cada vez que ele roda aumenta o x
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println("====================")
	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b()) // o x não é o mesmo para k e para b. Temos um closure

}

func soma(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	} //utiliza uma variavel fora do seu escopo- x está na func subjacente; outro escopo. x é único.
} 