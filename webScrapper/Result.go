package main

import "fmt"

//Result exported

type Result struct {
	typeOf 		string
	title    	string
	description string
}

func (r Result) String() string {
	return fmt.Sprint("[Type: ",r.typeOf, " - Title: ", r.title, " -Description: ", r.description," ]\n")
}