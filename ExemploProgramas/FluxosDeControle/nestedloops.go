import (
	"fmt"
)

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Print("hora: ", horas, "\n") //loop vai de 0 a 60
		for x := 0; x < 60; x++ {
			fmt.Print(horas, ":", x, "\n")

		} //para cada ciclo do loop externo, repito o interno 60 vezes.
	}

	//podemos também fazer com datas, meses, etc.

	fmt.println("============================")

}
