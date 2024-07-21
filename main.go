package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	res, err := http.Get("https://wikiwiki.jp/splatoon3mix/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
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
