import (
	"fmt"
)

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 25,
	}
	fmt.Println(x)
}
