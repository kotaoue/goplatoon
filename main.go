package main

import (
	"fmt"
	"log"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://wikiwiki.jp/splatoon3mix/"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	body, err := fetch(baseURL)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	doc.Find("div.navfold-container.clearfix").Each(func(i int, s *goquery.Selection) {
		group := s.Find("span.navfold-summary-label").Text()
		s.Find("li").Each(func(j int, li *goquery.Selection) {
			li.Find("a").Each(func(k int, a *goquery.Selection) {
				name := a.Text()
				link, exists := a.Attr("href")
				if exists {
					fmt.Printf("group = %s, name = %s, link = %s\n", group, name, link)
				}
			})
		})
	})
	return nil
}


func fetch(url string) (io.Reader, error){
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}