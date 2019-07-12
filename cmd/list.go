package cmd

import (
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Information descriptions.",
	Long:  `Information all branch descriptions.`,
	RunE:  execute,
}

func init() {
	listCmd.PersistentFlags().BoolP("only-list", "", false, "display list only")
	rootCmd.AddCommand(listCmd)
}

func execute(cmd *cobra.Command, args []string) error {
	configList, err := gitConfig.ConfigList()
	if err != nil {
		return err
	}

	onlyListFlg, _ := cmd.PersistentFlags().GetBool("only-list")
	t := tabby.New()
	if !onlyListFlg {
		t.AddHeader("BRANCH", "DESCRIPTION")
	}

	for _, info := range configList {
		t.AddLine(info.Branch, info.Description)
	}
	t.Print()
	return nil
}
