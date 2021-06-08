

import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)

//como fazer as duas funções rodarem de forma concorrente ?
var wg sync.WaitGroup

func main() {
	fmt.Println("Antes, tenho um total de: ", runtime.NumCPU(), "processadores")
	fmt.Println("Antes, tenho um total de: ", runtime.NumGoroutine(), "threads")
	//add(total de funções) - tenho 1 goroutine que meu programa vai ter que esperar.
	wg.Add(2)
	go func1()
	go func2()
	fmt.Println("Depois, tenho um total de:", runtime.NumGoroutine(), "threads")
	//espera
	wg.Wait()
}
func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1 ", i)
		//time.Sleep(20) //função vai dormir por 20ms  entre as execuções
	}
	//done!
	wg.Done()
}
func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2 ", i)
		//time.Sleep(20)
	}

	wg.Done()
}
