import (
	"fmt"
)

//mostre todos os anos desde o ano que você nasceu até o ano atual usando a sintaxe for { }
func main() {
	i := 2001
	for {
		if i > 2021 {
			break
		} //coloco uma condição dentro do loop em um if, e quando ela é mais atingida, saio do loop
		fmt.Printf("%v \n", i)
		i++

	}
}

