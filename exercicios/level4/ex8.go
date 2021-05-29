import (
	"fmt"
)

//map que tenha uma string como key, value como uma slice de strings representando, respectivamente o nome das pessoas e os seus Hobbies. Demostrar usando range
func main() {
	nomesHobbies := map[string][]string{
		"da Silva_Joana":    []string{"jogar basquete", "jogar Futebol"},
		"e Juiano_Henrique": []string{"jogar CS", "jogar Gartic"},
		"Manso_Carol":       []string{"fazer cursos", "ir a academia"},
	}
	fmt.Println(nomesHobbies)
	for key, value := range nomesHobbies {
		fmt.Println(key, ":")
		for i, v2 := range value {
			fmt.Println(i, "->", v2)
		}
		fmt.Println("\n================")

	}

}