
import (
	"fmt"
)

//mostre todos os anos desde o ano que você nasceu até o ano atual
func main() {
	i := 2001
	for i < 2021 {
		fmt.Printf("%v \n", i)
		i++
	}
}
