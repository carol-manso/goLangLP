package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

/*faz a requisição e transforma em documento*/
func scrap(url string) (r Result) {

	res, err := http.Get(url) //faz uma requisição de GET para o site desejado

	if err != nil {
		log.Fatal(err) //se tiver erro
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal("Status nao ideal", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body) //transforma o corpo da requisição em documento
	if err != nil {
		fmt.Println("ERROR: It can't parse html '", url, "'")
	}

	return scrapper(doc)
}

//navegar pelo documento
func scrapper(doc *goquery.Document) (r Result) {

	doc.Find("meta").Each(

		func(i int, s *goquery.Selection) {

			val, _ := s.Attr("property")

			if val == "og:title" {

				r.title, _ = s.Attr("content")

			} else if val == "og:site_name" {

				r.websiteName, _ = s.Attr("content")

			} else if val == "og:description" {

				r.description, _ = s.Attr("content")
			}
		},
	) //função anônima

	return

}
