
import (
	"fmt"
)

func main() {
	user := map[int]string{ //id, nome
		1:  "Joana",
		98: "Arthur",
		23: "Roberto",
		26: "Carlos",
	}
	fmt.Println(user)
	for key, value := range user { //nem sempre sai na ordem que inserimos. Range passa por todos os valores
		fmt.Println(key, value)
	}
	delete(user, 98)
	fmt.Println("\n", user)

}

