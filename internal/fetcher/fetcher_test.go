package fetcher

import (
	"strings"
	"testing"
)

func TestExtractStages(t *testing.T) {
	html := `<html><body>
<h4>ステージ</h4>
<div class="navfold-container">
  <span class="navfold-summary-label">ステージ一覧</span>
  <div class="navfold-content">
    <ul>
      <li><a href="/stage1">バンカラストリート</a></li>
      <li><a href="/stage2">ゴンズイ地区</a></li>
      <li><a href="/overview">概要</a></li>
      <li><a href="/area">面積あれこれ</a></li>
    </ul>
  </div>
</div>
</body></html>`

	stages, err := extractStages(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractStages returned error: %v", err)
	}

	expected := []string{"バンカラストリート", "ゴンズイ地区"}
	if len(stages) != len(expected) {
		t.Fatalf("expected %d stages, got %d: %v", len(expected), len(stages), stages)
	}
	for i, s := range expected {
		if stages[i] != s {
			t.Errorf("stages[%d]: expected %q, got %q", i, s, stages[i])
		}
	}
}

func TestExtractStages_Empty(t *testing.T) {
	html := `<html><body><p>no stages here</p></body></html>`

	stages, err := extractStages(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractStages returned error: %v", err)
	}
	if len(stages) != 0 {
		t.Errorf("expected empty stages, got %v", stages)
	}
}

func TestExtractSubWeapons(t *testing.T) {
	html := `<html><body>
<div class="navfold-container clearfix">
  <span class="navfold-summary-label">サブウェポン一覧</span>
  <div class="navfold-content">
    <ul>
      <li><a href="/sub1">スプラッシュボム</a></li>
      <li><a href="/sub2">キューバンボム</a></li>
      <li><a href="/sub">サブウェポン</a></li>
    </ul>
  </div>
</div>
</body></html>`

	subs, err := extractSubWeapons(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractSubWeapons returned error: %v", err)
	}

	expected := []string{"スプラッシュボム", "キューバンボム"}
	if len(subs) != len(expected) {
		t.Fatalf("expected %d sub weapons, got %d: %v", len(expected), len(subs), subs)
	}
	for i, s := range expected {
		if subs[i] != s {
			t.Errorf("subs[%d]: expected %q, got %q", i, s, subs[i])
		}
	}
}

func TestExtractSubWeapons_Empty(t *testing.T) {
	html := `<html><body><p>no sub weapons here</p></body></html>`

	subs, err := extractSubWeapons(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractSubWeapons returned error: %v", err)
	}
	if len(subs) != 0 {
		t.Errorf("expected empty sub weapons, got %v", subs)
	}
}

func TestExtractSpecialWeapons(t *testing.T) {
	html := `<html><body>
<div class="navfold-container clearfix">
  <span class="navfold-summary-label">スペシャルウェポン一覧</span>
  <div class="navfold-content">
    <ul>
      <li><a href="/sp1">ウルトラショット</a></li>
      <li><a href="/sp2">トリプルトルネード</a></li>
      <li><a href="/sp">スペシャルウェポン</a></li>
    </ul>
  </div>
</div>
</body></html>`

	specials, err := extractSpecialWeapons(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractSpecialWeapons returned error: %v", err)
	}

	expected := []string{"ウルトラショット", "トリプルトルネード"}
	if len(specials) != len(expected) {
		t.Fatalf("expected %d special weapons, got %d: %v", len(expected), len(specials), specials)
	}
	for i, s := range expected {
		if specials[i] != s {
			t.Errorf("specials[%d]: expected %q, got %q", i, s, specials[i])
		}
	}
}

func TestExtractSpecialWeapons_Empty(t *testing.T) {
	html := `<html><body><p>no special weapons here</p></body></html>`

	specials, err := extractSpecialWeapons(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractSpecialWeapons returned error: %v", err)
	}
	if len(specials) != 0 {
		t.Errorf("expected empty special weapons, got %v", specials)
	}
}

func TestExtractMainWeapons(t *testing.T) {
	html := `<html><body>
<div class="navfold-container clearfix">
  <span class="navfold-summary-label">シューター</span>
  <div class="navfold-content">
    <ul>
      <li><a href="/w1" title="ブキ/わかばシューター">わかばシューター</a></li>
      <li><a href="/w2" title="ブキ/シャープマーカー">シャープマーカー</a></li>
      <li><a href="/ov" title="ブキ/シューター属">シューター属</a></li>
      <li><a href="/no-title">タイトルなし</a></li>
    </ul>
  </div>
</div>
</body></html>`

	mains, err := extractMainWeapons(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractMainWeapons returned error: %v", err)
	}

	expected := []string{"わかばシューター", "シャープマーカー"}
	if len(mains) != len(expected) {
		t.Fatalf("expected %d main weapons, got %d: %v", len(expected), len(mains), mains)
	}
	for i, s := range expected {
		if mains[i] != s {
			t.Errorf("mains[%d]: expected %q, got %q", i, s, mains[i])
		}
	}
}

func TestExtractMainWeapons_Empty(t *testing.T) {
	html := `<html><body><p>no main weapons here</p></body></html>`

	mains, err := extractMainWeapons(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractMainWeapons returned error: %v", err)
	}
	if len(mains) != 0 {
		t.Errorf("expected empty main weapons, got %v", mains)
	}
}

func TestContains(t *testing.T) {
	targets := []string{"シューター一覧", "ローラー一覧", "チャージャー一覧"}

	tests := []struct {
		s    string
		want bool
	}{
		{"シューター", true},
		{"ローラー", true},
		{"スピナー", false},
		{"一覧", true},
	}

	for _, tt := range tests {
		got := contains(targets, tt.s)
		if got != tt.want {
			t.Errorf("contains(%v, %q) = %v, want %v", targets, tt.s, got, tt.want)
		}
	}
}

func TestExtractCategoryWeaponSpecs(t *testing.T) {
	html := `<html><body>
<table>
  <tr>
    <th>ブキ名</th>
    <th>サブウェポン</th>
    <th>スペシャルウェポン</th>
  </tr>
  <tr>
    <td>わかばシューター</td>
    <td>スプラッシュボム</td>
    <td>グレートバリア</td>
  </tr>
  <tr>
    <td>シャープマーカー</td>
    <td>キューバンボム</td>
    <td>ウルトラショット</td>
  </tr>
</table>
</body></html>`

	specs, err := extractCategoryWeaponSpecs(strings.NewReader(html), "シューター")
	if err != nil {
		t.Fatalf("extractCategoryWeaponSpecs returned error: %v", err)
	}

	if len(specs) != 2 {
		t.Fatalf("expected 2 specs, got %d: %v", len(specs), specs)
	}
	if specs[0].Name != "わかばシューター" {
		t.Errorf("specs[0].Name: expected %q, got %q", "わかばシューター", specs[0].Name)
	}
	if specs[0].Type != "シューター" {
		t.Errorf("specs[0].Type: expected %q, got %q", "シューター", specs[0].Type)
	}
	if specs[0].Sub != "スプラッシュボム" {
		t.Errorf("specs[0].Sub: expected %q, got %q", "スプラッシュボム", specs[0].Sub)
	}
	if specs[0].Special != "グレートバリア" {
		t.Errorf("specs[0].Special: expected %q, got %q", "グレートバリア", specs[0].Special)
	}
	if specs[1].Name != "シャープマーカー" {
		t.Errorf("specs[1].Name: expected %q, got %q", "シャープマーカー", specs[1].Name)
	}
}

func TestExtractCategoryWeaponSpecs_Empty(t *testing.T) {
	html := `<html><body><p>no weapons here</p></body></html>`

	specs, err := extractCategoryWeaponSpecs(strings.NewReader(html), "シューター")
	if err != nil {
		t.Fatalf("extractCategoryWeaponSpecs returned error: %v", err)
	}
	if len(specs) != 0 {
		t.Errorf("expected empty specs, got %v", specs)
	}
}

func TestExtractCategoryWeaponSpecs_NoNameColumn(t *testing.T) {
	// Table without a "ブキ名" header should be ignored
	html := `<html><body>
<table>
  <tr><th>その他</th><th>メモ</th></tr>
  <tr><td>foo</td><td>bar</td></tr>
</table>
</body></html>`

	specs, err := extractCategoryWeaponSpecs(strings.NewReader(html), "シューター")
	if err != nil {
		t.Fatalf("extractCategoryWeaponSpecs returned error: %v", err)
	}
	if len(specs) != 0 {
		t.Errorf("expected empty specs, got %v", specs)
	}
}

func TestExtractWeaponPerformance(t *testing.T) {
	html := `<html><body>
<table>
  <tr>
    <th>ブキ名</th>
    <th>重量区分</th>
    <th>射程</th>
    <th>発射レート</th>
  </tr>
  <tr>
    <td>わかばシューター</td>
    <td>軽量級</td>
    <td>中</td>
    <td>速い</td>
  </tr>
  <tr>
    <td>シャープマーカー</td>
    <td>軽量級</td>
    <td>長</td>
    <td>普通</td>
  </tr>
</table>
</body></html>`

	perf, err := extractWeaponPerformance(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractWeaponPerformance returned error: %v", err)
	}

	if len(perf) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(perf))
	}
	w := perf["わかばシューター"]
	if w.Weight != "軽量級" {
		t.Errorf("Weight: expected %q, got %q", "軽量級", w.Weight)
	}
	if w.Range != "中" {
		t.Errorf("Range: expected %q, got %q", "中", w.Range)
	}
	if w.FireRate != "速い" {
		t.Errorf("FireRate: expected %q, got %q", "速い", w.FireRate)
	}
}

func TestExtractWeaponPerformance_Empty(t *testing.T) {
	html := `<html><body><p>no data here</p></body></html>`

	perf, err := extractWeaponPerformance(strings.NewReader(html))
	if err != nil {
		t.Fatalf("extractWeaponPerformance returned error: %v", err)
	}
	if len(perf) != 0 {
		t.Errorf("expected empty map, got %v", perf)
	}
}

func TestWeaponSpecString(t *testing.T) {
	spec := WeaponSpec{
		Name:     "わかばシューター",
		Type:     "シューター",
		Sub:      "スプラッシュボム",
		Special:  "グレートバリア",
		Weight:   "軽量級",
		Range:    "中",
		FireRate: "速い",
	}
	got := spec.String()
	if !strings.Contains(got, "わかばシューター") {
		t.Errorf("String() missing Name: %q", got)
	}
	if !strings.Contains(got, "シューター") {
		t.Errorf("String() missing Type: %q", got)
	}
	if !strings.Contains(got, "軽量級") {
		t.Errorf("String() missing Weight: %q", got)
	}
}

func TestBuildWeaponURL(t *testing.T) {
	tests := []struct {
		href string
		want string
	}{
		{"https://example.com/weapon", "https://example.com/weapon"},
		{"/splatoon3mix/weapon", "https://wikiwiki.jp/splatoon3mix/weapon"},
		{"weapon", BaseURL + "weapon"},
	}
	for _, tt := range tests {
		got := buildWeaponURL(tt.href)
		if got != tt.want {
			t.Errorf("buildWeaponURL(%q) = %q, want %q", tt.href, got, tt.want)
		}
	}
}
