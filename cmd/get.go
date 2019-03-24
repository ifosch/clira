package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ifosch/clira/pkg"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get issue details",
	Long: `
This subcommand allows getting an issue details.

example: clira get DVX-567
`,
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := clira.GetClient()
		if err != nil {
			log.Fatal(err)
		}
		issue, _, err := jiraClient.Issue.Get(args[0], nil)

		fmt.Printf("%s:    %+v\n", issue.Key, issue.Fields.Summary)
		fmt.Printf("Type:     %s\n", issue.Fields.Type.Name)
		fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
		fmt.Printf("Reporter: %s\n", issue.Fields.Reporter.Name)
		fmt.Printf("Assignee: %s\n", issue.Fields.Assignee.Name)
		fmt.Printf("Status:   %s\n", issue.Fields.Status.Name)
		fmt.Printf("\nDescription: %s\n", issue.Fields.Description)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
