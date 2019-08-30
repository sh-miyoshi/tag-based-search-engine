package cmd

import (
	"fmt"
	"github.com/sh-miyoshi/tag-based-search-engine/pkg/engine"
	"github.com/spf13/cobra"
)

var deleteFile string
var deleteTags []string
var deleteTargetFile string

func init() {
	deleteFileCmd.Flags().StringVarP(&deleteFile, "name", "n", "", "a name of delete file")
	deleteFileCmd.MarkFlagRequired("name")

	deleteTagCmd.Flags().StringVarP(&deleteTargetFile, "name", "n", "", "a name of delete file")
	deleteTagCmd.Flags().StringSliceVar(&deleteTags, "tags", []string{}, "a list of tags to delete")
	deleteTagCmd.MarkFlagRequired("name")
	deleteTagCmd.MarkFlagRequired("tags")

	deleteCmd.AddCommand(deleteTagCmd)
	deleteCmd.AddCommand(deleteFileCmd)

	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete file or tags",
	Long:  `Delete file or tags`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("delete command requires subcommand")
	},
}

var deleteFileCmd = &cobra.Command{
	Use:   "file",
	Short: "delete file to database",
	Long:  `delete file to database`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.DeleteFile(deleteFile, dbFile); err != nil {
			fmt.Printf("Failed to delete file Error: %v", err)
			return
		}
		fmt.Println("Succeeded")
	},
}

var deleteTagCmd = &cobra.Command{
	Use:   "tags",
	Short: "delete tags of the file",
	Long:  `delete tags of the file`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := engine.DeleteTag(deleteTargetFile, deleteTags, dbFile); err != nil {
			fmt.Printf("Failed to delete tags Error: %v", err)
			return
		}
		fmt.Println("Succeeded")
	},
}
