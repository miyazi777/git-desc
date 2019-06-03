package cmd

import (
	"fmt"

	"github.com/InVisionApp/tabular"
	"github.com/mattn/go-runewidth"
	"github.com/miyazi777/git-desc/git"
	"github.com/miyazi777/git-desc/terminal"
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

		width, err := terminal.GetWidth()
		branchWidth := int(float32(width) * 0.3)
		descWidth := int(float32(width) * 0.7)

		tab := tabular.New()
		tab.Col("branch", "Branch", branchWidth)
		tab.Col("desc", "Description", descWidth)
		format := tab.Print("*")

		for branchName, description := range descriptionMap {
			fmt.Printf(format, runewidth.Truncate(branchName, branchWidth, "..."), runewidth.Truncate(description, descWidth, "..."))
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
