package main

import (
	"fmt"
	"os"

	"github.com/kotaoue/goplatoon/internal/fetcher"
	"github.com/spf13/cobra"
)

var (
	target string
)

var rootCmd = &cobra.Command{
	Use:   "goplatoon",
	Short: "Splatoon 3 data fetcher",
	Long:  "A CLI tool to fetch Splatoon 3 data from wikiwiki.jp/splatoon3mix",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch target {
		case "main":
			return fetchAndPrint(fetcher.FetchMainWeapons)
		case "sub":
			return fetchAndPrint(fetcher.FetchSubWeapons)
		case "sp":
			return fetchAndPrint(fetcher.FetchSpecialWeapons)
		case "stage":
			return fetchAndPrint(fetcher.FetchStages)
		default:
			return fmt.Errorf("invalid target: %s (must be 'main', 'sub', 'sp' or 'stage')", target)
		}
	},
}

var mainCmd = &cobra.Command{
	Use:   "main",
	Short: "Fetch main weapons",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fetchAndPrint(fetcher.FetchMainWeapons)
	},
}

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Fetch sub weapons",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fetchAndPrint(fetcher.FetchSubWeapons)
	},
}

var spCmd = &cobra.Command{
	Use:   "sp",
	Short: "Fetch special weapons",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fetchAndPrint(fetcher.FetchSpecialWeapons)
	},
}

var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Fetch stages",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fetchAndPrint(fetcher.FetchStages)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "main", "fetch target: main, sub, sp or stage")
	rootCmd.AddCommand(mainCmd)
	rootCmd.AddCommand(subCmd)
	rootCmd.AddCommand(spCmd)
	rootCmd.AddCommand(stageCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
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
