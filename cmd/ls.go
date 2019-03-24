package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ifosch/clira/pkg"
)

// lsCmd represents the get command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Retrieve a list of issues",
	Long: `
This subcommand allows getting issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := clira.GetClient()
		if err != nil {
			log.Fatal(err)
		}
		results, _, err := jiraClient.Issue.Search("", nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(
			"KEY\tTYPE\tSTATUS\tDESCRIPTION\n",
		)
		for _, result := range results {
			fmt.Printf(
				"%s\t%s\t%v\t%s\n",
				result.Key,
				result.Fields.Type.Name,
				result.Fields.Status.Name,
				result.Fields.Summary,
			)
		}
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
