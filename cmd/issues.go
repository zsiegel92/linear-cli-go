package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/zach/linear_cli_go/actions"
	"github.com/zach/linear_cli_go/linear"
	"github.com/zach/linear_cli_go/mock"
	"github.com/zach/linear_cli_go/models"
	"github.com/zach/linear_cli_go/ui"
)

const (
	DefaultIssueLimit = 10
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
