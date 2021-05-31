import (
	"fmt"
)

//fazer uma struct anonima com um map e um slice.
func main() {
	anonimo := struct {
		listaSacolao map[int]string
		extras       []string
	}{
		listaSacolao: map[int]string{
			1: "banana",
			2: "maçã",
			3: "laranja",
			4: "alface",
		},
		extras: []string{"abóbora", "biscoito recheado", "sucrilhos", "nutela"},
	}
	fmt.Println(anonimo)
}
