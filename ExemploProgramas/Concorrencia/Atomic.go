package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//solução para as raceConditions
func main() {
	var contador int64
	contador = 0
	totaldegoroutines := 15

	var wg sync.WaitGroup

	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1) //1= delta- muda 1 no endereço
			runtime.Gosched()
			fmt.Println("Contador: \t", atomic.LoadInt64(&contador))
			wg.Done()

		}()
		//fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("Valor final:", contador)

}
