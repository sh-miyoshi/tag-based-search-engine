package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register file to database",
	Long:  `Register file to database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Succeeded")
	},
}
