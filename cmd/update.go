package cmd

import (
	"fmt"
	"log"

	"github.com/ifosch/clira/pkg"
	"github.com/spf13/cobra"
)

// udpateCmd represents the get command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Issue status",
	Long: `
This subcommand allows updating an issue status

example: clira update DVX-567 resolved
`,
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := clira.GetClient()
		if err != nil {
			log.Fatal(err)
		}
		issue, _, err := jiraClient.Issue.Get(args[0], nil)
		if err != nil {
			log.Fatal(err)
		}

		transitions, _, err := jiraClient.Issue.GetTransitions(issue.ID)
		if err != nil {
			log.Fatal(err)
		}

		found, err := clira.GetTransition(transitions, args[1])
		if err != nil {
			log.Fatal(err)
		}
		_, err = jiraClient.Issue.DoTransition(issue.ID, found.ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("desired state: %+v\n", found)
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
