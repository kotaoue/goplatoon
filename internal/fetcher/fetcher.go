package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	BaseURL = "https://wikiwiki.jp/splatoon3mix/"
)

func Fetch(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res.Body, nil
}

func FetchStages() ([]string, error) {
	body, err := Fetch(BaseURL)
	if err != nil {
		return nil, err
	}

	return extractStages(body)
}

func extractStages(reader io.Reader) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var stages []string

	doc.Find("div.navfold-container.clearfix").Each(func(i int, s *goquery.Selection) {
		group := s.Find("span.navfold-summary-label").Text()

		if contains([]string{"ステージ一覧"}, group) {
			s.Find("li").Each(func(j int, li *goquery.Selection) {
				li.Find("a").Each(func(k int, a *goquery.Selection) {
					name := a.Text()
					_, exists := a.Attr("href")
					if exists {
						stages = append(stages, name)
					}
				})
			})
		}
	})
	return stages, nil
}

func contains(targets []string, s string) bool {
	for _, target := range targets {
		if strings.Contains(target, s) {
			return true
		}
	}
	return false
}
