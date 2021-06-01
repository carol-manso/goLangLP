//crie uma função que receba um parametro variadico do tipo int e retorne a soma deles
//passe um slice of ints como argumento
//faça uma função que receba um slice of int como argumento
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(soma(slice...))
	fmt.Println(soma2(slice))

}

func soma(x ...int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total
}
func soma2(x []int) int {
	total := 0
	for _, value := range x {
		total += value
	}
	return total

}