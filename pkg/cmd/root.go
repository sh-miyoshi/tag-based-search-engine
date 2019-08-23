package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tsearch",
	Short: "tsearch is a command line tool for tag-based-search-engine",
	Long:  "tsearch is a command line tool for tag-based-search-engine",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute method run root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
