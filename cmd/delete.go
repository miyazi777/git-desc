package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the info command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete current branch descrpition.",
	Long:  "Delete current branch descrpition.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		err = description.DeleteDescription()
		if err != nil {
			return err
		}

		err = page.DeletePage()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
