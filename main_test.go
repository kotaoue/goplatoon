package main

import (
	"errors"
	"testing"
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
