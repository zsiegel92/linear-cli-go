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
	// Default number of issues to fetch
	DefaultIssueLimit = 10
)

// RunIssueSelector runs the issue selector flow
func RunIssueSelector(useMock bool, onlyMine bool, projectID string) {
	var issues []models.LinearIssue
	var err error

	if useMock {
		if onlyMine {
			// For mock data, simulate showing only Alice's issues
			issues = mock.GetMockIssuesForUser("Alice")
		} else {
			issues = mock.GetMockIssues()
		}
	} else {
		client := linear.NewClient("") // Will use LINEAR_API_KEY from env
		issues, err = client.GetIssues(onlyMine, projectID, DefaultIssueLimit)
		if err != nil {
			log.Fatalf("Error fetching issues from Linear: %v", err)
		}
	}

	if len(issues) == 0 {
		fmt.Println("No issues found")
		return
	}

	// Select an issue
	selectedIssue, err := ui.SelectIssue(issues)
	if err != nil {
		log.Fatalf("Error selecting issue: %v", err)
	}

	if selectedIssue == nil {
		fmt.Println("No issue selected")
		os.Exit(0)
	}

	// Show the selected issue
	fmt.Printf("\n✅ Selected issue: %s\n\n", selectedIssue.Title)

	// Select an action
	selectedAction, err := ui.SelectAction(selectedIssue)
	if err != nil {
		log.Fatalf("Error selecting action: %v", err)
	}

	// Execute the action
	if err := actions.ExecuteAction(selectedAction, selectedIssue); err != nil {
		log.Fatalf("Error executing action: %v", err)
	}
}
