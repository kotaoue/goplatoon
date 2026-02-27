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
