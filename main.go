package main

import (
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
	stages, err := fetcher.FetchStages()
	if err != nil {
		return err
	}

	for _, stage := range stages {
		fmt.Println(stage)
	}

	return nil
}
