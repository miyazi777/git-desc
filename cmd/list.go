package cmd

import (
	"fmt"

	"gitdesc/git"
	"github.com/InVisionApp/tabular"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Information descriptions.",
	Long:  `Information all branch descriptions.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		descriptionMap, err := git.BuildDescriptionMap()
		if err != nil {
			return err
		}

		tab := tabular.New()
		tab.Col("branch", "Branch", 30)
		tab.Col("desc", "Description", 40)
		format := tab.Print("*")

		for branchName, description := range descriptionMap {
			fmt.Printf(format, branchName, description)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
