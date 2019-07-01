package cmd

import (
	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pageSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Page set command.",
	Long:  "Page set command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var text string
		text, _ = cmd.PersistentFlags().GetString("page")
		if text == "" {
			page, err := branch.Page()
			if err != nil {
				return err
			}

			text, err = shell.EditTextByEditor(page)
			if err != nil {
				return err
			}
		}

		err := branch.SetPage(text)
		return err
	},
}

func init() {
	pageSetCmd.PersistentFlags().StringP("page", "m", "", "page url")
	pageCmd.AddCommand(pageSetCmd)
}
