
import (
	"fmt"
)
func main() {
	amigos := map[string]int{
		"alfedo": 554852, //key :valor,
		"joana":  998856,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["joana"]) //valor de amigos na posição joana
	//adicionando valores:
	amigos["gopher"] = 856943
	fmt.Println(amigos["gopher"])
	será, ok := amigos["Inês|istente"] // tem numero 0. ok= false-> chave inesistente não existe no mapa. 0 porque não existe
	if !ok {
		fmt.Println("não existe!")
	} else {
		fmt.Println(será)
	}
	

	//fmt.Println(será,ok)


}