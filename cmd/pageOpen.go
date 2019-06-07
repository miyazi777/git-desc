package cmd

import (
	"os/exec"

	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pageOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Page open command.",
	Long:  "Page open command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		branchName, err := shell.GetCurrentBranch()
		if err != nil {
			return err
		}

		page, err := shell.GetPage(branchName)
		if err != nil {
			return err
		}

		command := exec.Command("open", page)
		cmdErr := command.Run()
		if cmdErr != nil {
			return cmdErr
		}

		return nil
	},
}

func init() {
	pageCmd.AddCommand(pageOpenCmd)
}
