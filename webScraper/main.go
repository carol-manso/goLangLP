package main

import (
	"fmt"
	"time"
)

func main() {
	urlToProcess := []string{
		"https://stackoverflow.com/questions/214741/what-is-a-stackoverflowerror#:~:text=A%20StackOverflowError%20is%20a%20runtime,excessive%20deep%20or%20infinite%20recursion.&text=In%20the%20above%20case%2C%20it,avoided%20by%20doing%20programmatic%20changes.",
		"https://github.com/PuerkitoBio/goquery",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
	}

	fmt.Println("Without goroutines:")
	fmt.Println()

	ini := time.Now()

	for _, url := range urlToProcess {

		r := scrap(url)
		fmt.Println(r)

	}
	tempo1 := time.Since(ini).Seconds()
	fmt.Println("==========================\n(Took ",
		tempo1, "secs)\n==========================")
	fmt.Println()

	ini = time.Now()

	r := make(chan Result)

	go scrapListURL(urlToProcess, r)

	fmt.Println("With goroutines:")
	fmt.Println()

	for url := range r {
		fmt.Println(url)
	}
	tempo2 := time.Since(ini).Seconds()
	fmt.Println("==========================\n(Took ",
		tempo2, "secs)\n==========================")

	speedup := tempo1 / tempo2
	fmt.Println("Speed Up:", speedup)
}
