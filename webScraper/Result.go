package main

import "fmt"

//Necessário primeira letra maiúscula para exportar do arquivo
type Result struct {
	websiteName string
	title       string
	description string
}

func (r Result) String() string {

	return fmt.Sprint("\nWebsite name: ",
		r.websiteName, " -\nTitle: ", r.title, " -\nDescription:\n", r.description,
		"\n----------------------------------------------- \n")

}
