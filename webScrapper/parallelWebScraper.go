package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

 func scrapListURL(urlToProcess []string, rchan chan Result) {

 	defer close(rchan)

 	var results = []chan Result{}

 	for i, url := range urlToProcess {

 		results = append(results, make(chan Result))
 		go scrapParallel(url, results[i])

 	}

 	for i := range results {
 		for r1 := range results[i] {
 			rchan <- r1
 		}
 	}

 }

 func scrapParallel(url string, rchan chan Result) {

 	defer close(rchan)

 	res, err := http.Get(url)

 	if err != nil {

 		fmt.Println("ERROR: It can't scrap '", url, "'")

 	}

 	defer res.Body.Close()
 	
	if(res.StatusCode != 200){

		log.Fatal("not ideal Status",res.StatusCode)

	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

 	if err != nil {

 		fmt.Println("ERROR: It can't parse html '", url, "'")

 	}
 	
 	rchan <- scrapper(doc)
 }