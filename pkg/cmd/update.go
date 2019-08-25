package cmd

import (
	"fmt"
	"github.com/sh-miyoshi/tag-based-search-engine/pkg/engine"
	"github.com/spf13/cobra"
)

var updateTags []string
var updateTargetFile string

func init() {
	assignCmd.Flags().StringVarP(&updateTargetFile, "name", "n", "", "a name of register file")
	assignCmd.Flags().StringSliceVar(&updateTags, "tags", []string{}, "a list of tags for assign/remove")
	assignCmd.MarkFlagRequired("name")
	assignCmd.MarkFlagRequired("tags")

	rootCmd.AddCommand(assignCmd)
}

var assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign tags of the file",
	Long:  `Assign tags of the file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.Assign(updateTargetFile, updateTags, dbFile); err != nil {
			fmt.Printf("Failed to assign tags Error: %v", err)
			return
		}
		fmt.Println("Succeeded")
	},
}
