
import (
	"fmt"
	"math"
)

// tipo quadrado, tipo circulo -> método área para cada um
//tipo figura- interface- qualquer tipo que tem o método área
//função info - recebe um tipo de fig e retorna sua área
//valor para circulo e para quadrado
//usar info para demostrar a área de ambos

type circulo struct {
	raio int
}
type quadrado struct {
	lado int
}

func (c circulo) calcularArea() {
	area := (float64(c.raio) * float64(c.raio) * math.Pi)
	fmt.Println("Área do círculo: ", area)
}
func (q quadrado) calcularArea() {
	fmt.Println("Área do quadrado: ", q.lado*q.lado)
}

type figura interface {
	calcularArea()
}

func info(f figura) {
	f.calcularArea()
}

func main() {
	quadrado1 := quadrado{4}
	circulo1 := circulo{3}
	info(quadrado1)
	info(circulo1)
}
