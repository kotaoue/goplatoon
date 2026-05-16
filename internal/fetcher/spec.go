package fetcher

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	WikiBaseURL = "https://wikiwiki.jp"
)

// WeaponSpec holds detailed specifications of a main weapon.
type WeaponSpec struct {
	Name     string // ブキ名
	Type     string // ブキ種類
	Sub      string // サブウェポン
	Special  string // スペシャルウェポン
	Weight   string // 重量区分
	Range    string // 射程
	FireRate string // 発射レート
}

func (s WeaponSpec) String() string {
	return fmt.Sprintf("名前: %s\n種類: %s\nサブ: %s\nスペシャル: %s\n重量区分: %s\n射程: %s\n発射レート: %s",
		s.Name, s.Type, s.Sub, s.Special, s.Weight, s.Range, s.FireRate)
}

type weaponEntry struct {
	Name string
	Type string
	HRef string
}

// FetchMainWeaponSpecs fetches detailed specifications for all main weapons.
func FetchMainWeaponSpecs() ([]WeaponSpec, error) {
	body, err := Fetch(BaseURL)
	if err != nil {
		return nil, err
	}

	entries, err := extractMainWeaponEntries(body)
	if err != nil {
		return nil, err
	}

	var specs []WeaponSpec
	for _, entry := range entries {
		weaponURL := buildWeaponURL(entry.HRef)
		wbody, err := Fetch(weaponURL)
		if err != nil {
			continue
		}
		spec, err := extractWeaponSpec(wbody)
		if err != nil {
			continue
		}
		spec.Name = entry.Name
		spec.Type = entry.Type
		specs = append(specs, spec)
	}
	return specs, nil
}

func buildWeaponURL(href string) string {
	if strings.HasPrefix(href, "http") {
		return href
	}
	if strings.HasPrefix(href, "/") {
		return WikiBaseURL + href
	}
	return BaseURL + href
}

func extractMainWeaponEntries(reader io.Reader) ([]weaponEntry, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var entries []weaponEntry
	weaponTypes := []string{"シューター", "ローラー", "チャージャー", "スロッシャー", "スピナー", "マニューバー", "シェルター", "ブラスター", "フデ", "ストリンガー", "ワイパー"}

	doc.Find("div.navfold-container.clearfix").Each(func(i int, s *goquery.Selection) {
		label := s.Find("span.navfold-summary-label").Text()

		for _, weaponType := range weaponTypes {
			if strings.Contains(label, weaponType) {
				s.Find("div.navfold-content li a").Each(func(j int, a *goquery.Selection) {
					title, exists := a.Attr("title")
					if exists && strings.HasPrefix(title, "ブキ/") {
						weaponName := strings.TrimPrefix(title, "ブキ/")
						if !strings.HasSuffix(weaponName, "属") {
							href, _ := a.Attr("href")
							entries = append(entries, weaponEntry{
								Name: weaponName,
								Type: weaponType,
								HRef: href,
							})
						}
					}
				})
				break
			}
		}
	})
	return entries, nil
}

func extractWeaponSpec(reader io.Reader) (WeaponSpec, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return WeaponSpec{}, err
	}

	spec := WeaponSpec{}

	doc.Find("table tr").Each(func(i int, row *goquery.Selection) {
		th := strings.TrimSpace(row.Find("th").Text())
		td := strings.TrimSpace(row.Find("td").First().Text())

		switch {
		case strings.Contains(th, "サブウェポン"):
			spec.Sub = td
		case strings.Contains(th, "スペシャルウェポン"):
			spec.Special = td
		case strings.Contains(th, "重量区分"):
			spec.Weight = td
		case strings.Contains(th, "射程"):
			spec.Range = td
		case strings.Contains(th, "発射レート"):
			spec.FireRate = td
		}
	})

	return spec, nil
}
