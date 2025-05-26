package ui

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/zach/linear_cli_go/models"
	"github.com/zach/linear_cli_go/utils"
)

func DisplayIssue(issue models.LinearIssue) string {
	return utils.FormatIssueDisplay(issue)
}

func previewIssue(issue models.LinearIssue, teamKeys []string) string {
	var preview strings.Builder
	
	projectName := "<No Project Specified For Issue>"
	projectSlug := ""
	if issue.Project != nil {
		projectName = issue.Project.Name
		projectSlug = utils.GetSlug(issue.Project.Name)
	}
	
	teamColor := utils.GetTeamColor(issue.Team.Key, teamKeys)
	projectLine := utils.BoldUnderlineText(projectName)
	if projectSlug != "" {
		projectLine += " - " + utils.ColorText(teamColor, fmt.Sprintf("(%s)", projectSlug))
	}
	preview.WriteString(projectLine + "\n")
	
	titleLine := utils.ColorText(utils.Blue, utils.BoldText(issue.Title))
	if issue.Estimate != nil {
		titleLine += fmt.Sprintf(" - (%s)", utils.FormatEstimate(issue.Estimate))
	}
	preview.WriteString(titleLine + "\n")
	
	if issue.Creator.DisplayName != "" {
		createdAt, err := time.Parse(time.RFC3339, issue.CreatedAt)
		if err == nil {
			preview.WriteString(fmt.Sprintf("Created by %s %s\n", issue.Creator.DisplayName, createdAt.Format("2006-01-02 15:04:05")))
		}
	}
	
	if issue.UpdatedAt != "" {
		preview.WriteString(fmt.Sprintf("Updated %s\n", utils.FormatTimeAgo(issue.UpdatedAt)))
	}
	
	if issue.BranchName != "" {
		preview.WriteString(utils.BoldText(issue.BranchName) + "\n")
	}
	
	if issue.URL != "" {
		preview.WriteString(utils.BoldText(issue.URL) + "\n")
	}
	
	preview.WriteString("\n")
	
	if issue.Description != "" {
		preview.WriteString(issue.Description)
	}
	
	return preview.String()
}

func SelectIssue(issues []models.LinearIssue) (*models.LinearIssue, error) {
	if len(issues) == 0 {
		return nil, fmt.Errorf("no issues to select from")
	}


	teamKeys := utils.GetUniqueTeamKeys(issues)
	
	index, err := fuzzyfinder.Find(
		issues,
		func(i int) string {
			return utils.FormatIssueDisplay(issues[i])
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return previewIssue(issues[i], teamKeys)
		}),
	)
	if err != nil {
		if err.Error() == "^C" || err.Error() == "^D" {
			fmt.Println("\n👋 Goodbye!")
			os.Exit(0)
		}
		return nil, fmt.Errorf("selection failed: %w", err)
	}

	return &issues[index], nil
}

type ActionItem struct {
	Label       string
	Action      models.Action
	Description string
}

func SelectAction(issue *models.LinearIssue) (models.Action, error) {
	actions := []ActionItem{
		{
			Label:       "📋 Copy issue URL",
			Action:      models.CopyIssueURL,
			Description: "Copy the Linear issue URL to clipboard",
		},
		{
			Label:       "🌿 Copy branch name",
			Action:      models.CopyBranchName,
			Description: "Copy the branch name to clipboard",
		},
		{
			Label:       "🌐 Open in browser",
			Action:      models.OpenInBrowser,
			Description: "Open the issue in your default browser",
		},
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
			var preview strings.Builder
			preview.WriteString(utils.BoldText(action.Label) + "\n\n")
			preview.WriteString(action.Description + "\n\n")
			
				switch action.Action {
			case models.CopyBranchName:
				preview.WriteString(fmt.Sprintf("Branch name: %s", utils.BoldText(issue.BranchName)))
			case models.CopyIssueURL, models.OpenInBrowser:
				preview.WriteString(fmt.Sprintf("URL: %s", utils.BoldText(issue.URL)))
			}
			
			return preview.String()
		}),
	)
	if err != nil {
		if err.Error() == "^C" || err.Error() == "^D" {
			fmt.Println("\n👋 Goodbye!")
			os.Exit(0)
		}
		return "", fmt.Errorf("action selection failed: %w", err)
	}

	return actions[index].Action, nil
}
