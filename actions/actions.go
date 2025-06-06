package actions

import (
	"fmt"
	"strings"

	"github.com/zsiegel92/linear-cli-go/models"
	"github.com/zsiegel92/linear-cli-go/utils"
)

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

func GenerateBranchName(issue *models.LinearIssue) string {
	title := strings.ToLower(issue.Title)
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, "_", "-")
	
	var builder strings.Builder
	for _, char := range title {
		if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '-' {
			builder.WriteRune(char)
		}
	}
	title = builder.String()
	
	for strings.Contains(title, "--") {
		title = strings.ReplaceAll(title, "--", "-")
	}
	title = strings.Trim(title, "-")
	
	if len(title) > 50 {
		title = title[:50]
		title = strings.Trim(title, "-")
	}
	
	branchName := fmt.Sprintf("%s-%s", strings.ToLower(issue.Team.Key), title)
	
	return branchName
}