// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/ifosch/clira"
	"github.com/spf13/cobra"
)

// searchIssuesCmd represents the search_issues command
var searchIssuesCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches issues using JQL",
	Long:  `Searches issues using JQL.`,
	Run: func(cmd *cobra.Command, args []string) {
		clira.Login()
		comments, _ := cmd.Flags().GetBool("comments")
		pattern, _ := cmd.Flags().GetString("pattern")
		clira.SearchIssues(args, comments, pattern)
	},
}

func init() {
	issuesCmd.AddCommand(searchIssuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// search_issuesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// search_issuesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchIssuesCmd.Flags().BoolP("comments", "c", false, "Show comments")
	searchIssuesCmd.Flags().StringP("pattern", "", "", "Regular expression to hook comments")
}
