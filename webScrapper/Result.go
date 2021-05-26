package main

import "fmt"

//Necessário primeira letra maiúscula para exportar do arquivo
type Result struct {
	typeOf      string
	title       string
	description string
}

func (r Result) String() string {

	return fmt.Sprint("\nType: ", 
						r.typeOf, " -\nTitle: ", r.title, " -\nDescription:\n", r.description, 
					  "\n----------------------------------------------- \n")

}
