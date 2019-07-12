package cmd

import (
	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

var command shell.Command

// infoCmd represents the info command
var pageOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Page open command.",
	Long:  "Page open command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		page, err := page.Get()
		if err != nil {
			return err
		}

		_, err = command.Run("open", page)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	pageCmd.AddCommand(pageOpenCmd)
	command = shell.CommandImpl{}
}
