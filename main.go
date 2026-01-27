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
	mode := flag.String("mode", "main", "fetch mode: main, sub, sp or stage")
	flag.Parse()

	switch *mode {
	case "main":
		return fetchAndPrint(fetcher.FetchMainWeapons)
	case "sub":
		return fetchAndPrint(fetcher.FetchSubWeapons)
	case "sp":
		return fetchAndPrint(fetcher.FetchSpecialWeapons)
	case "stage":
		return fetchAndPrint(fetcher.FetchStages)
	default:
		return fmt.Errorf("invalid mode: %s (must be 'main', 'sub', 'sp' or 'stage')", *mode)
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
