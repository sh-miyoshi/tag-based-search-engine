package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update tags of the file",
	Long:  `Update tags of the file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Succeeded")
	},
}
