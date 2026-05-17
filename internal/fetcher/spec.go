package fetcher

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	WikiBaseURL    = "https://wikiwiki.jp"
	weaponPerfPath = "ブキ/ブキ性能"
)

// weaponTypeCategories lists each weapon type and its wiki category page path.
var weaponTypeCategories = []struct {
	Type string
	Path string
}{
	{"シューター", "ブキ/シューター属"},
	{"ローラー", "ブキ/ローラー属"},
	{"チャージャー", "ブキ/チャージャー属"},
	{"スロッシャー", "ブキ/スロッシャー属"},
	{"スピナー", "ブキ/スピナー属"},
	{"マニューバー", "ブキ/マニューバー属"},
	{"シェルター", "ブキ/シェルター属"},
	{"ブラスター", "ブキ/ブラスター属"},
	{"フデ", "ブキ/フデ属"},
	{"ストリンガー", "ブキ/ストリンガー属"},
	{"ワイパー", "ブキ/ワイパー属"},
}

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

// FetchMainWeaponSpecs fetches detailed specifications for all main weapons.
// Sub/special are read from per-type category pages (e.g. ブキ/シューター属);
// weight, range, and fire-rate are read from the single performance page (ブキ/ブキ性能).
func FetchMainWeaponSpecs() ([]WeaponSpec, error) {
	var specs []WeaponSpec
	nameIndex := make(map[string]int)

	for _, cat := range weaponTypeCategories {
		body, err := Fetch(BaseURL + cat.Path)
		if err != nil {
			continue
		}
		catSpecs, err := extractCategoryWeaponSpecs(body, cat.Type)
		if err != nil {
			continue
		}
		for _, s := range catSpecs {
			if _, exists := nameIndex[s.Name]; !exists {
				nameIndex[s.Name] = len(specs)
				specs = append(specs, s)
			}
		}
	}

	body, err := Fetch(BaseURL + weaponPerfPath)
	if err == nil {
		perfMap, err := extractWeaponPerformance(body)
		if err == nil {
			for name, perf := range perfMap {
				if idx, ok := nameIndex[name]; ok {
					specs[idx].Weight = perf.Weight
					specs[idx].Range = perf.Range
					specs[idx].FireRate = perf.FireRate
				}
			}
		}
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

// extractCategoryWeaponSpecs parses a weapon category page (e.g. シューター属) and
// returns specs with Name, Type, Sub, and Special populated from the main weapon table.
func extractCategoryWeaponSpecs(reader io.Reader, weaponType string) ([]WeaponSpec, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	var specs []WeaponSpec

	doc.Find("table").EachWithBreak(func(_ int, table *goquery.Selection) bool {
		headers := tableHeaders(table)
		nameIdx := sliceIndexOf(headers, "ブキ名")
		subIdx := sliceIndexOfContains(headers, "サブウェポン")
		spIdx := sliceIndexOfContains(headers, "スペシャルウェポン")

		if nameIdx < 0 {
			return true // not a weapon table; try next
		}

		table.Find("tr").Each(func(k int, row *goquery.Selection) {
			if k == 0 {
				return // skip header row
			}
			cells := row.Find("td")
			if cells.Length() <= nameIdx {
				return
			}
			name := strings.TrimSpace(cells.Eq(nameIdx).Text())
			if name == "" {
				return
			}
			spec := WeaponSpec{Name: name, Type: weaponType}
			if subIdx >= 0 && cells.Length() > subIdx {
				spec.Sub = strings.TrimSpace(cells.Eq(subIdx).Text())
			}
			if spIdx >= 0 && cells.Length() > spIdx {
				spec.Special = strings.TrimSpace(cells.Eq(spIdx).Text())
			}
			specs = append(specs, spec)
		})

		return false // stop at first matching table
	})

	return specs, nil
}

// extractWeaponPerformance parses the ブキ性能 page and returns a map of
// weapon name → partial WeaponSpec containing Weight, Range, and FireRate.
func extractWeaponPerformance(reader io.Reader) (map[string]WeaponSpec, error) {
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	result := make(map[string]WeaponSpec)

	doc.Find("table").EachWithBreak(func(_ int, table *goquery.Selection) bool {
		headers := tableHeaders(table)
		nameIdx := sliceIndexOf(headers, "ブキ名")
		weightIdx := sliceIndexOfContains(headers, "重量区分")
		rangeIdx := sliceIndexOfContains(headers, "射程")
		rateIdx := sliceIndexOfContains(headers, "発射レート")

		if nameIdx < 0 || (weightIdx < 0 && rangeIdx < 0 && rateIdx < 0) {
			return true // not the performance table; try next
		}

		table.Find("tr").Each(func(k int, row *goquery.Selection) {
			if k == 0 {
				return // skip header row
			}
			cells := row.Find("td")
			if cells.Length() <= nameIdx {
				return
			}
			name := strings.TrimSpace(cells.Eq(nameIdx).Text())
			if name == "" {
				return
			}
			spec := WeaponSpec{}
			if weightIdx >= 0 && cells.Length() > weightIdx {
				spec.Weight = strings.TrimSpace(cells.Eq(weightIdx).Text())
			}
			if rangeIdx >= 0 && cells.Length() > rangeIdx {
				spec.Range = strings.TrimSpace(cells.Eq(rangeIdx).Text())
			}
			if rateIdx >= 0 && cells.Length() > rateIdx {
				spec.FireRate = strings.TrimSpace(cells.Eq(rateIdx).Text())
			}
			result[name] = spec
		})

		return false // stop at first matching table
	})

	return result, nil
}

// tableHeaders returns the trimmed text of all th cells in the first row of a table.
func tableHeaders(table *goquery.Selection) []string {
	var headers []string
	table.Find("tr").First().Find("th").Each(func(_ int, th *goquery.Selection) {
		headers = append(headers, strings.TrimSpace(th.Text()))
	})
	return headers
}

// sliceIndexOf returns the index of the first element equal to target, or -1.
func sliceIndexOf(slice []string, target string) int {
	for i, s := range slice {
		if s == target {
			return i
		}
	}
	return -1
}

// sliceIndexOfContains returns the index of the first element containing target, or -1.
func sliceIndexOfContains(slice []string, target string) int {
	for i, s := range slice {
		if strings.Contains(s, target) {
			return i
		}
	}
	return -1
}
