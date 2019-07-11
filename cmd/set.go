package cmd

import (
	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current branch description.",
	Long:  `Set current branch description.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		var text string
		text, _ = cmd.PersistentFlags().GetString("message")
		if text == "" {
			desc, err := description.Get()
			if err != nil {
				return err
			}

			text, err = shell.EditTextByEditor(desc)
			if err != nil {
				return err
			}
		}

		err = description.Set(text)
		return err
	},
}

func init() {
	setCmd.PersistentFlags().StringP("message", "m", "", "description message")
	rootCmd.AddCommand(setCmd)
}
