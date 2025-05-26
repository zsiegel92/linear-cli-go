package ui

import (
	"fmt"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/zach/linear_cli_go/models"
	"github.com/zach/linear_cli_go/utils"
)

// SelectIssueWithFzf displays a list of issues with side-by-side preview using fuzzyfinder
func SelectIssueWithFzf(issues []models.LinearIssue) (*models.LinearIssue, error) {
	if len(issues) == 0 {
		return nil, fmt.Errorf("no issues to select from")
	}

	index, err := fuzzyfinder.Find(
		issues,
		func(i int) string {
			return utils.FormatIssueDisplay(issues[i])
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			issue := issues[i]
			
			preview := fmt.Sprintf("Title: %s\n", issue.Title)
			preview += fmt.Sprintf("Team: %s\n", issue.Team.Key)
			
			if issue.Project != nil {
				preview += fmt.Sprintf("Project: %s\n", issue.Project.Name)
			}
			
			if issue.Assignee != nil {
				preview += fmt.Sprintf("Assignee: %s\n", issue.Assignee.DisplayName)
			} else {
				preview += "Assignee: UNASSIGNED\n"
			}
			
			preview += fmt.Sprintf("Status: %s\n", issue.State.Name)
			
			if issue.PriorityLabel != "" {
				preview += fmt.Sprintf("Priority: %s\n", issue.PriorityLabel)
			}
			
			if issue.Estimate != nil {
				preview += fmt.Sprintf("Estimate: %s points\n", utils.FormatEstimate(issue.Estimate))
			}
			
			if issue.UpdatedAt != "" {
				preview += fmt.Sprintf("Updated: %s\n", utils.FormatTimeAgo(issue.UpdatedAt))
			}
			
			preview += fmt.Sprintf("URL: %s\n", issue.URL)
			
			if issue.BranchName != "" {
				preview += fmt.Sprintf("Branch: %s\n", issue.BranchName)
			}
			
			if issue.Description != "" {
				preview += fmt.Sprintf("\nDescription:\n%s", issue.Description)
			}
			
			return preview
		}),
		fuzzyfinder.WithPromptString("Select an issue > "),
	)
	
	if err != nil {
		// Fallback to original promptui selector if fzf fails
		fmt.Println("Falling back to basic selector...")
		return SelectIssue(issues)
	}

	return &issues[index], nil
}

// SelectActionWithFzf displays action options using fuzzyfinder
func SelectActionWithFzf(issue *models.LinearIssue) (models.Action, error) {
	actions := []struct {
		Label  string
		Action models.Action
		Desc   string
	}{
		{"📋 Copy issue URL", models.CopyIssueURL, "Copy the Linear issue URL to clipboard"},
		{"🌿 Copy branch name", models.CopyBranchName, "Copy the branch name to clipboard"},
		{"🌐 Open in browser", models.OpenInBrowser, "Open the issue in your default browser"},
	}

	index, err := fuzzyfinder.Find(
		actions,
		func(i int) string {
			return actions[i].Label
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			action := actions[i]
			
			preview := fmt.Sprintf("Action: %s\n", action.Label)
			preview += fmt.Sprintf("Description: %s\n\n", action.Desc)
			preview += fmt.Sprintf("Issue: [%s] %s\n", issue.Team.Key, issue.Title)
			
			return preview
		}),
		fuzzyfinder.WithPromptString("What would you like to do? > "),
	)
	
	if err != nil {
		// Fallback to original promptui selector if fzf fails
		fmt.Println("Falling back to basic action selector...")
		return SelectAction(issue)
	}

	return actions[index].Action, nil
}