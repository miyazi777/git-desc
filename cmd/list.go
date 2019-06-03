package cmd

import (
	"fmt"

	"github.com/InVisionApp/tabular"
	"github.com/mattn/go-runewidth"
	"github.com/miyazi777/git-desc/git"
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
		tab.Col("branch", "Branch", 40)
		tab.Col("desc", "Description", 50)
		format := tab.Print("*")

		for branchName, description := range descriptionMap {
			fmt.Printf(format, runewidth.Truncate(branchName, 40, "..."), runewidth.Truncate(description, 50, "..."))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
