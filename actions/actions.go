package actions

import (
	"fmt"
	"strings"

	"github.com/zach/linear_cli_go/models"
	"github.com/zach/linear_cli_go/utils"
)

// ExecuteAction performs the specified action on the given issue
func ExecuteAction(action models.Action, issue *models.LinearIssue) error {
	switch action {
	case models.CopyIssueURL:
		if err := utils.CopyToClipboard(issue.URL); err != nil {
			return fmt.Errorf("failed to copy URL to clipboard: %w", err)
		}
		fmt.Printf("✅ Copied issue URL to clipboard: %s\n", issue.URL)
		return nil

	case models.CopyBranchName:
		branchName := issue.BranchName
		if branchName == "" {
			// Generate a branch name if none exists
			branchName = GenerateBranchName(issue)
		}
		if err := utils.CopyToClipboard(branchName); err != nil {
			return fmt.Errorf("failed to copy branch name to clipboard: %w", err)
		}
		fmt.Printf("✅ Copied branch name to clipboard: %s\n", branchName)
		return nil

	case models.OpenInBrowser:
		if err := utils.OpenInBrowser(issue.URL); err != nil {
			return fmt.Errorf("failed to open issue in browser: %w", err)
		}
		fmt.Printf("✅ Opened issue in browser: %s\n", issue.URL)
		return nil

	default:
		return fmt.Errorf("unknown action: %s", action)
	}
}

// GenerateBranchName creates a branch name from the issue
func GenerateBranchName(issue *models.LinearIssue) string {
	// Convert title to kebab-case
	title := strings.ToLower(issue.Title)
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "_", "-")
	
	// Remove special characters
	var builder strings.Builder
	for _, char := range title {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			builder.WriteRune(char)
		}
	}
	title = builder.String()
	
	// Remove consecutive dashes and trim
	for strings.Contains(title, "--") {
		title = strings.ReplaceAll(title, "--", "-")
	}
	title = strings.Trim(title, "-")
	
	// Limit length
	if len(title) > 50 {
		title = title[:50]
		title = strings.Trim(title, "-")
	}
	
	// Combine with team key
	branchName := fmt.Sprintf("%s-%s", strings.ToLower(issue.Team.Key), title)
	
	return branchName
}