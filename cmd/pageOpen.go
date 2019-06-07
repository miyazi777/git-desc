package cmd

import (
	"fmt"

	//"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pageOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Page open command.",
	Long:  "Page open command.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("page open command")
		return nil
		//branchName, err := shell.GetCurrentBranch()
		//if err != nil {
		//	return err
		//}
		//description, err := shell.GetDesctiption(branchName)
		//if err != nil {
		//	return err
		//}
		//fmt.Printf("description: %s\n", description)
		//return nil
	},
}

func init() {
	pageCmd.AddCommand(pageOpenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
