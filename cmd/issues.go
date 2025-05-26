package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/zsiegel92/linear-cli-go/actions"
	"github.com/zsiegel92/linear-cli-go/linear"
	"github.com/zsiegel92/linear-cli-go/mock"
	"github.com/zsiegel92/linear-cli-go/models"
	"github.com/zsiegel92/linear-cli-go/ui"
)

const (
	DefaultIssueLimit = 80
)

func RunIssueSelector(useMock bool, onlyMine bool, projectID string) {
	var issues []models.LinearIssue
	var err error

	if useMock {
		if onlyMine {
			issues = mock.GetMockIssuesForUser("Alice")
		} else {
			issues = mock.GetMockIssues()
		}
	} else {
		client := linear.NewClient("")
		issues, err = client.GetIssues(onlyMine, projectID, DefaultIssueLimit)
		if err != nil {
			log.Fatalf("Error fetching issues from Linear: %v", err)
		}
	}

	if len(issues) == 0 {
		fmt.Println("No issues found")
		return
	}

	selectedIssue, err := ui.SelectIssue(issues)
	if err != nil {
		log.Fatalf("Error selecting issue: %v", err)
	}

	if selectedIssue == nil {
		fmt.Println("No issue selected")
		os.Exit(0)
	}

	fmt.Printf("\n✅ Selected issue: %s\n\n", selectedIssue.Title)

	selectedAction, err := ui.SelectAction(selectedIssue)
	if err != nil {
		log.Fatalf("Error selecting action: %v", err)
	}

	if err := actions.ExecuteAction(selectedAction, selectedIssue); err != nil {
		log.Fatalf("Error executing action: %v", err)
	}
}
