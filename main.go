package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

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

	return extract(body, []string{"ステージ一覧"})
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

func extract(reader io.Reader, targets []string) error {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return err
	}

	doc.Find("div.navfold-container.clearfix").Each(func(i int, s *goquery.Selection) {
		group := s.Find("span.navfold-summary-label").Text()

		if contains(targets, group) {
			s.Find("li").Each(func(j int, li *goquery.Selection) {
				li.Find("a").Each(func(k int, a *goquery.Selection) {
					name := a.Text()
					_, exists := a.Attr("href")
					if exists {
						fmt.Printf("%s\n", name)
					}
				})
			})
			}
	})
	return nil
}

func contains(targets []string, s string) bool {
	for _, target := range targets {
		if strings.Contains(target, s) {
			return true
		}
	}
	return false
}
