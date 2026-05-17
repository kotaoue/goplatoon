package main

import (
	"errors"
	"testing"

	"github.com/kotaoue/goplatoon/internal/fetcher"
)

func TestFetchAndPrint_Success(t *testing.T) {
	fetch := func() ([]string, error) {
		return []string{"itemA", "itemB"}, nil
	}
	if err := fetchAndPrint(fetch); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestFetchAndPrint_Empty(t *testing.T) {
	fetch := func() ([]string, error) {
		return []string{}, nil
	}
	if err := fetchAndPrint(fetch); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestFetchAndPrint_Error(t *testing.T) {
	want := errors.New("fetch failed")
	fetch := func() ([]string, error) {
		return nil, want
	}
	if err := fetchAndPrint(fetch); err != want {
		t.Errorf("expected error %v, got %v", want, err)
	}
}

func TestFetchAndPrintSpecs_Success(t *testing.T) {
	fetch := func() ([]fetcher.WeaponSpec, error) {
		return []fetcher.WeaponSpec{
			{Name: "わかばシューター", Type: "シューター"},
			{Name: "シャープマーカー", Type: "シューター"},
		}, nil
	}
	if err := fetchAndPrintSpecs(fetch); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestFetchAndPrintSpecs_Empty(t *testing.T) {
	fetch := func() ([]fetcher.WeaponSpec, error) {
		return []fetcher.WeaponSpec{}, nil
	}
	if err := fetchAndPrintSpecs(fetch); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestFetchAndPrintSpecs_Error(t *testing.T) {
	want := errors.New("fetch failed")
	fetch := func() ([]fetcher.WeaponSpec, error) {
		return nil, want
	}
	if err := fetchAndPrintSpecs(fetch); err != want {
		t.Errorf("expected error %v, got %v", want, err)
	}
}
