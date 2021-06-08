package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	go func() {
		canal <- 42 //coloco uma info no meu canal
	}()

	fmt.Println(<-canal) //goroutine querendo retirar info do canal.
}
