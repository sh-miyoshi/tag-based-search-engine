package cmd

import (
	"fmt"
	"github.com/sh-miyoshi/tag-based-search-engine/pkg/engine"
	"github.com/spf13/cobra"
)

var registerFile string
var updateTags []string
var updateTargetFile string

func init() {
	registerFileCmd.Flags().StringVarP(&registerFile, "name", "n", "", "a name of register file")
	registerFileCmd.MarkFlagRequired("name")
	registerTagCmd.Flags().StringVarP(&updateTargetFile, "name", "n", "", "a name of register file")
	registerTagCmd.Flags().StringSliceVar(&updateTags, "tags", []string{}, "a list of tags for assign/remove")
	registerTagCmd.MarkFlagRequired("name")
	registerTagCmd.MarkFlagRequired("tags")

	registerCmd.AddCommand(registerTagCmd)
	registerCmd.AddCommand(registerFileCmd)

	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register file or tags to database",
	Long:  `Register file or tags to database`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("register command requires subcommand")
	},
}

var registerFileCmd = &cobra.Command{
	Use:   "file",
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

var registerTagCmd = &cobra.Command{
	Use:   "tags",
	Short: "Register tags of the file",
	Long:  `Register tags of the file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.AssignTag(updateTargetFile, updateTags, dbFile); err != nil {
			fmt.Printf("Failed to assign tags Error: %v", err)
			return
		}
		fmt.Println("Succeeded")
	},
}
