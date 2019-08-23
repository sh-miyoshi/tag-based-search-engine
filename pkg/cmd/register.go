package cmd

import (
	"fmt"
	"github.com/sh-miyoshi/tag-based-search-engine/pkg/engine"
	"github.com/spf13/cobra"
)

var registerFile string

func init() {
	registerCmd.Flags().StringVarP(&registerFile, "name", "n", "", "a name of register file")
	registerCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register file to database",
	Long:  `Register file to database`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.Register(registerFile, dbFile); err != nil {
			fmt.Printf("Failed to register file: %s Error: %v", registerFile, err)
			return
		}
		fmt.Println("Succeeded")
	},
}
