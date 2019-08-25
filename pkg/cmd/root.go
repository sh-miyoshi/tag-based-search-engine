package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var dbFile string

func init() {
	const defaultDBFile = "database.json"

	rootCmd.PersistentFlags().StringVar(&dbFile, "database", defaultDBFile, "The file path of Database")
}

var rootCmd = &cobra.Command{
	Use:   "tsearch",
	Short: "tsearch is a command line tool for tag-based-search-engine",
	Long:  "tsearch is a command line tool for tag-based-search-engine",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute method run root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
