package main

import (
	"fmt"
	"time"
)



func say(str string) {

	for i:=0; i<5 ;i++ {

		fmt.Println(str)
		time.Sleep(100*time.Millisecond)
		
		
 	}
	
}


func main (){
	
	go say("Hello")
	say("World")
	
	

	
	
	
	

	
	
	

}

