

import (
	"fmt"
)

//usando o map do ex anterior, remover uma entrada e mostrar o map inteiro com range
func main() {
	nomesHobbies := map[string][]string{
		"da Silva_Joana":    []string{"jogar basquete", "jogar Futebol"},
		"e Juiano_Henrique": []string{"jogar CS", "jogar Gartic"},
		"Manso_Carol":       []string{"fazer cursos", "ir a academia"},
	}
	delete(nomesHobbies, "e Juiano_Henrique")
	fmt.Println(nomesHobbies)
	fmt.Println("")
	for key, value := range nomesHobbies {
		fmt.Println(key, ":")
		for i, v2 := range value {
			fmt.Println(i, "->", v2)
		}
		fmt.Println("\n================")

	}

}
