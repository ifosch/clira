package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/ifosch/clira/pkg"
)

type newSprint struct {
	Name          string `json:"name"`
	OriginBoardID int    `json:"originBoardId"`
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new sprint",
	Long: `
This subcommand allows sprint creation.`,
	Run: func(cmd *cobra.Command, args []string) {
		jiraClient, err := clira.GetClient()
		if err != nil {
			log.Fatal(err)
		}

		board, err := clira.GetBoard(jiraClient)
		if err != nil {
			log.Fatal(err)
		}

		createdSprint := &newSprint{
			Name:          args[0],
			OriginBoardID: board.ID,
		}
		req, err := jiraClient.NewRequest(
			"POST",
			"rest/agile/1.0/sprint",
			createdSprint,
		)
		if err != nil {
			log.Fatal(err)
		}
		result := new(newSprint)
		_, err = jiraClient.Do(req, result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New sprint created: ", result.Name)
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
