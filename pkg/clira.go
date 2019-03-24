package clira

import (
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
