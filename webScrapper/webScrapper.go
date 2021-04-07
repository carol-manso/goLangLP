package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)



func scrap(url string) (r Result) {
	
	res,err := http.Get(url)
	
	if err!=nil{
		log.Fatal(err)
	}

	defer res.Body.Close()
	
	if(res.StatusCode != 200){
		log.Fatal("Status nao ideal",res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		 	fmt.Println("ERROR: It can't parse html '", url, "'")
		 }
	
	
	
	

	return scrapper(doc)
}

func scrapper(doc *goquery.Document)(r Result){

	doc.Find("meta").Each(func (i int, s *goquery.Selection)  {

		val,_:= s.Attr("property")
		
			if  val == "og:site_name"{
				
				r.title,_ = s.Attr("content")

			}else if val == "og:type"{

				r.typeOf,_ = s.Attr("content")
				
			}else if val =="og:description"{

				r.description,_ = s.Attr("content")
			}
	

		})

		return

}
	



