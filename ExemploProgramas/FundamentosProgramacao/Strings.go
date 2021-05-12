

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	S := `Helloooo, 
				Play
						Groooooound` //formatada da forma que quisermos.
	sb := []byte(s)
	fmt.Printf("%v,%T\n", s, s)
	fmt.Printf("%v,%T\n", S, S)
	fmt.Printf("%v,%T\n", sb, sb) //valor de cada byte de cada caractere
	for _, v := range sb {
		fmt.Printf("\n%v- %T- %#U- %#x", v, v, v, v)
	}
	//me dá carctere por caractere. UTF-8 não dá problema aqui
	fmt.Println("\n=====================================")

	for i := 0; i < len(s); i++ {
		fmt.Printf("\n%v- %T- %#U- %#x", s[i], s[i], s[i], s[i])
	} //mesmo resultado, porém, o for loop me dá as coisas em bytes e caracteres acentuados ou especiais que ocupam mais de 1 byte saem meui doidões por sair cada byte deles.

	/*\n%v- %T- %#U- %#x"
	%v= valor numerico
	%T= tipo
	%#U-> unicode code point.
	%#x = valor em hexadecimal.
	*/

}
