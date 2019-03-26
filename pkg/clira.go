package clira

import (
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

// GetClient returns a logged client using BasicAuth.
func GetClient() (client *jira.Client, err error) {
	base := os.Getenv("JIRA_HOST")
	tp := jira.BasicAuthTransport{
		Username: os.Getenv("JIRA_USER"),
		Password: os.Getenv("JIRA_PASSWORD"),
	}

	client, err = jira.NewClient(tp.Client(), base)
	return
}

func GetTransition(
	transitions []jira.Transition,
	transitionName string,
) (found jira.Transition, err error) {
	transitionNames := ""

	for _, transition := range transitions {
		transitionNames = transitionNames + ", " + transition.Name
		if transition.Name == transitionName {
			found = transition
		}
	}

	if found.Name == "" {
		return found, fmt.Errorf("Unknown transition, valid transtions are: %s \n", transitionNames[1:])
	}
	return
}

type boardResults struct {
	MaxResults int          `json:"maxResults"`
	StartAt    int          `json:"startAt"`
	IsLast     bool         `json:"isLast"`
	Values     []jira.Board `json:"values"`
}

// GetBoard returns the board for a user. If there is more than a board for a user, it returns an error. If the user can't see any board, it returns an error.
func GetBoard(client *jira.Client) (board *jira.Board, err error) {
	req, err := client.NewRequest(
		"GET",
		"rest/agile/1.0/board",
		nil,
	)
	if err != nil {
		return
	}
	result := new(boardResults)
	_, err = client.Do(req, result)
	if err != nil {
		return
	}
	if len(result.Values) < 1 {
		return nil, fmt.Errorf("Too few boards")
	}
	if len(result.Values) > 1 {
		return nil, fmt.Errorf("Too many boards")
	}
	board = &result.Values[0]
	return
}
