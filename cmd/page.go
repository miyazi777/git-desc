package cmd

import (
	//"fmt"

	//"github.com/miyazi777/git-desc/shell"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var pageCmd = &cobra.Command{
	Use:   "page",
	Short: "Page command.",
	Long:  "Page command.",
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	fmt.Println("page command")
	//	return nil
	//	//branchName, err := shell.GetCurrentBranch()
	//	//if err != nil {
	//	//	return err
	//	//}
	//	//description, err := shell.GetDesctiption(branchName)
	//	//if err != nil {
	//	//	return err
	//	//}
	//	//fmt.Printf("description: %s\n", description)
	//	//return nil
	//},
}

func init() {
	rootCmd.AddCommand(pageCmd)
}
