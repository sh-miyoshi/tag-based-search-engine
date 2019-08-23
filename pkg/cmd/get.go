package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get files from tags",
	Long:  `Get files from tags`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Succeeded")
	},
}
