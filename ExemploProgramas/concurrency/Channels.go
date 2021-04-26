package main

import "fmt"



func exemplo(c chan string){
	defer close(c)
	strs:=[3]string{"limao","banana","maçã"}

	for _,i := range strs{

		c <-i
	}

}

func main(){

	channel :=make(chan string)

	go exemplo(channel)

	for i :=range(channel){
		fmt.Println(i)
	}

}