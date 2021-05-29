
import (
	"fmt"
)

func main() {
	sabores := []string{"calabresa", "margherita", "quatro queijos", "a moda", "espinafre"}
	//excluindo um item de um slice -> seleciono tudo antes e tudo depois e faço um merge com o append
	sabores = append(sabores[:2], sabores[3:]...) //crio uma slice nova que vai ser o merge
	fmt.Println(sabores)
}
