import (
	"fmt"
)

// programa para demostrar switch com statement sendo uma str e id esporteFavorito
func main() {
	esporteFavorito := "basquete"
	switch esporteFavorito {
	case "basquete":
		fmt.Println("Gostei da escolha")
	case "curling":
		fmt.Println("Gente, quero muito praticar um dia para ver como é")
	case "futebol":
		fmt.Println("Para qual time torce ?")
	}
}
