package main

import (
	"fmt"
)

//anexar a slice 52; depois adc 53,54,55 usando uma declaração; depois demonstrar e adc mais uma slice
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	for _, value := range x {
		fmt.Println("->", value)
	}

}
