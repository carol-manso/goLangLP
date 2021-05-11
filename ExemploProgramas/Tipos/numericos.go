

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	c := "∑" //3 runes- todos com 32 bits por caracter, mas o tanto de bytes usados é diferente.
	fmt.Printf("%v,%v,%v\n", a, b, c)
	d := []byte(a) //converter algo em bytes.
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v,%v,%v", d, e, f) //quantos bytes usamos.
	x := 10
	y := 10.0
	fmt.Printf("\n%v, %T\n", y, y)
	fmt.Printf("\n%v, %T\n", x, x)
	k := runtime.GOOS //sistema operacional
	l := runtime.GOARCH
	fmt.Println("Seu computadorm de SO = " + k + " possui arquitetura de " + l + " bits")
	var i uint16
	i = 65535          //número limite de um uint16
	fmt.Println(i + 1) //overflow

}
