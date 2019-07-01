package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Information current branch descrpition.",
	Long:  "Information current branch descrpition.",
	RunE: func(cmd *cobra.Command, args []string) error {
		description, err := branch.Description()
		if err != nil {
			return err
		}

		page, err := branch.Page()
		if err != nil {
			return err
		}

		fmt.Printf("Description: %s\n", description)
		fmt.Printf("       Page: %s\n", page)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
