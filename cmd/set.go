package cmd

import (
	"github.com/spf13/cobra"

	"github.com/miyazi777/git-desc/shell"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current branch description.",
	Long:  `Set current branch description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		branchName, err := shell.GetCurrentBranch()
		if err != nil {
			return err
		}

		var text string
		text, _ = cmd.PersistentFlags().GetString("message")
		if text == "" {
			branchName, err := shell.GetCurrentBranch()
			if err != nil {
				return err
			}

			description, err := shell.GetDesctiption(branchName)
			if err != nil {
				return err
			}

			text, err = shell.EditTextByEditor(description)
			if err != nil {
				return err
			}
		}

		err = shell.SetDescription(branchName, text)
		return err
	},
}

func init() {
	setCmd.PersistentFlags().StringP("message", "m", "", "description message")
	rootCmd.AddCommand(setCmd)
}
