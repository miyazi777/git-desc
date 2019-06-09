package cmd

import (
	"github.com/miyazi777/git-desc/git"
	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pageOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Page open command.",
	Long:  "Page open command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var branch git.Branch

		page, err := branch.Page()
		if err != nil {
			return err
		}

		_, err = shell.Run("open", page)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	pageCmd.AddCommand(pageOpenCmd)
}
