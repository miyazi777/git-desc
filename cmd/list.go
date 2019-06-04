package cmd

import (
	"errors"
	"fmt"

	"github.com/InVisionApp/tabular"
	"github.com/mattn/go-runewidth"
	"github.com/miyazi777/git-desc/git"
	"github.com/miyazi777/git-desc/terminal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Information descriptions.",
	Long:  `Information all branch descriptions.`,
	RunE:  exec,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func exec(cmd *cobra.Command, args []string) error {
	descriptionMap, err := git.BuildDescriptionMap()
	if err != nil {
		return err
	}

	width, err := terminal.GetWidth()

	branchWidthPer, err := getBranchWidthPer()
	if err != nil {
		return err
	}

	descWidthPer, err := getDescWidthPer()
	if err != nil {
		return err
	}

	branchWidth := int(float64(width) * branchWidthPer)
	descWidth := int(float64(width) * descWidthPer)

	tab := tabular.New()
	tab.Col("branch", "Branch", branchWidth)
	tab.Col("desc", "Description", descWidth)
	format := tab.Print("*")

	for branchName, description := range descriptionMap {
		fmt.Printf(format, runewidth.Truncate(branchName, branchWidth, "..."), runewidth.Truncate(description, descWidth, "..."))
	}
	return nil
}

func getBranchWidthPer() (float64, error) {
	checkVal := viper.Get("list.branch.width")
	if checkVal == nil {
		return 0, errors.New("Require error: list.branch.width")
	}

	val := viper.GetFloat64("list.branch.width")
	if val < 0.0 || val > 1.00 {
		return 0, errors.New("Range error: list.branch.width")
	}
	return val, nil
}

func getDescWidthPer() (float64, error) {
	checkVal := viper.Get("list.description.width")
	if checkVal == nil {
		return 0, errors.New("Require error: list.description.width")
	}

	val := viper.GetFloat64("list.description.width")
	if val < 0.0 || val > 1.00 {
		return 0, errors.New("Range error: list.description.width")
	}
	return val, nil
}
