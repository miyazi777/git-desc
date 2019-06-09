package cmd

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Page command.",
	Long:  "Page command.",
}

func init() {
	rootCmd.AddCommand(pageCmd)
}
