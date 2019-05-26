package cmd

import (
	"errors"
	"github.com/spf13/cobra"

	"gitdesc/git"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current branch description.",
	Long:  `Set current branch description.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires description")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		branchName, err := git.GetCurrentBranch()
		if err != nil {
			return err
		}
		git.SetDescription(branchName, args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
