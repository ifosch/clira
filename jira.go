// Copyright © 2016 Ignasi Fosch
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

package clira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"regexp"
)

// JIRAClient is the interface to JIRA API.
var JIRAClient *jira.Client

// SprintCommentsForIssue prints all comments for an issue.
func SprintCommentsForIssue(issue jira.Issue, pattern string) (string, error) {
	issueData, _, err := JIRAClient.Issue.Get(issue.ID)
	var commentsSelected []*jira.Comment
	for _, comment := range issueData.Fields.Comments.Comments {
		if pattern == "" {
			commentsSelected = append(commentsSelected, comment)
		} else {
			match, _ := regexp.MatchString(pattern, comment.Body)
			if match {
				commentsSelected = append(commentsSelected, comment)
			}
		}
	}
	output := ""
	for _, comment := range commentsSelected {
		commentText := ""
		if pattern == "" {
			commentText = comment.Body
		} else {
			match, _ := regexp.MatchString(pattern, comment.Body)
			if match {
				commentText = comment.Body
			}
		}
		if commentText != "" {
			output = fmt.Sprintf("%s\n\t- %s", output, commentText)
		}
	}
	return output, err
}

// SearchIssues searches issues matching a JQL query.
func SearchIssues(args []string, comments bool, pattern string) {
	jql := ""
	if len(args) >= 1 {
		jql = args[0]
	}
	issues, _, err := JIRAClient.Issue.Search(jql, nil)
	if err != nil {
		panic(fmt.Errorf("Error getting issues for '%v': %v\n", jql, err))
	}
	output := ""
	for _, issue := range issues {
		output = fmt.Sprintf("%s%s %s", output, issue.Key, issue.Fields.Summary)
		if comments {
			commentsOutput, err := SprintCommentsForIssue(issue, pattern)
			if err != nil {
				panic(fmt.Errorf("Error getting issue %s: %v\n", issue.Key, err))
			}
			output = fmt.Sprintf("%s%s", output, commentsOutput)
		}
		output = fmt.Sprintf("%s\n", output)
	}
	fmt.Print(output)
}

// ListProjects lists all projects from JIRA client.
func ListProjects() {
	projects, _, err := JIRAClient.Project.GetList()
	if err != nil {
		panic(fmt.Errorf("Error getting projects %v\n", err))
	}
	for _, project := range *projects {
		fmt.Printf("%s: %s\n", project.Key, project.Name)
	}
}

// Login initializes JIRA client.
func Login() {
	fmt.Println("Logging in")
	config := getConfigFromViper()
	var err error
	JIRAClient, err = jira.NewClient(nil, config.URL)
	if err != nil {
		panic(fmt.Errorf("Error logging in: %v\n", err))
	}
	_, err = JIRAClient.Authentication.AcquireSessionCookie(
		config.Username,
		config.Password)
	if err != nil {
		panic(fmt.Errorf("Error authenticating in: %v\n", err))
	}
}
