package cmd

import (
	"fmt"
	"github.com/sh-miyoshi/tag-based-search-engine/pkg/engine"
	"github.com/spf13/cobra"
)

var getTags []string

func init() {
	getCmd.Flags().StringSliceVar(&getTags, "tags", []string{}, "a list of tags to filter data")

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get files from tags",
	Long:  `Get files from tags`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := engine.Get(getTags, dbFile)
		if err != nil {
			fmt.Printf("Failed to get data, Error: %v", err)
			return
		}

		for _, d := range res {
			fmt.Println(d)
		}
	},
}
