package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kotaoue/goplatoon/internal/fetcher"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	mode := flag.String("mode", "stage", "fetch mode: stage, sub or sp")
	flag.Parse()

	switch *mode {
	case "stage":
		return fetchAndPrint(fetcher.FetchStages)
	case "sub":
		return fetchAndPrint(fetcher.FetchSubWeapons)
	case "sp":
		return fetchAndPrint(fetcher.FetchSpecialWeapons)
	default:
		return fmt.Errorf("invalid mode: %s (must be 'stage', 'sub' or 'sp')", *mode)
	}
}

func fetchAndPrint(fetchFunc func() ([]string, error)) error {
	items, err := fetchFunc()
	if err != nil {
		return err
	}

	for _, item := range items {
		fmt.Println(item)
	}

	return nil
}
