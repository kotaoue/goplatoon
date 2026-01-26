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

	doc.Find("h4").Each(func(i int, h4 *goquery.Selection) {
		if strings.Contains(h4.Text(), "ステージ") {
			navfold := h4.Next()
			if navfold.HasClass("navfold-container") {
				label := navfold.Find("span.navfold-summary-label").Text()
				if strings.Contains(label, "ステージ一覧") {
					navfold.Find("div.navfold-content li a").Each(func(j int, a *goquery.Selection) {
						name := a.Text()
						if name != "概要" && name != "面積あれこれ"{
							stages = append(stages, name)
						}
					})
				}
			}
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
