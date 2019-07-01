package cmd

import (
	"errors"
	"fmt"

	"github.com/InVisionApp/tabular"
	"github.com/mattn/go-runewidth"
	"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Information descriptions.",
	Long:  `Information all branch descriptions.`,
	RunE:  execute,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func execute(cmd *cobra.Command, args []string) error {
	descriptionMap, err := branch.DescriptionMap()
	if err != nil {
		return err
	}

	width, err := shell.GetWidth()

	branchWidthPer, err := getPercentConfig("list.branch.width")
	if err != nil {
		return err
	}

	descWidthPer, err := getPercentConfig("list.description.width")
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

func getPercentConfig(key string) (float64, error) {
	checkVal := viper.Get(key)
	if checkVal == nil {
		return 0, errors.New("Require error: " + key)
	}

	val := viper.GetFloat64(key)
	if val < 0.0 || val > 1.00 {
		return 0, errors.New("Range error: " + key)
	}
	return val, nil
}
