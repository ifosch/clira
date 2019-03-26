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
