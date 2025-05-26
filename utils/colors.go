package utils

import (
	"github.com/zach/linear_cli_go/models"
)

// ANSI color codes
const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	
	// Colors
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	Gray    = "\033[90m"
)

// Color functions
func ColorText(color, text string) string {
	return color + text + Reset
}

func BoldText(text string) string {
	return Bold + text + Reset
}

func UnderlineText(text string) string {
	return Underline + text + Reset
}

func BoldUnderlineText(text string) string {
	return Bold + Underline + text + Reset
}

// Secondary colors for team identification (matching TypeScript version)
var SecondaryColors = []string{Yellow, Cyan, Magenta}

// GetTeamColor returns a consistent color for a team key
func GetTeamColor(teamKey string, teamKeys []string) string {
	for i, key := range teamKeys {
		if key == teamKey {
			return SecondaryColors[i%len(SecondaryColors)]
		}
	}
	return Reset // fallback
}

// GetUniqueTeamKeys extracts unique team keys from issues
func GetUniqueTeamKeys(issues []models.LinearIssue) []string {
	seen := make(map[string]bool)
	var teamKeys []string
	
	for _, issue := range issues {
		if !seen[issue.Team.Key] {
			seen[issue.Team.Key] = true
			teamKeys = append(teamKeys, issue.Team.Key)
		}
	}
	
	return teamKeys
}

// FormatIssueDisplayWithColor formats issue for display with colors (matching TypeScript style)
// This version returns plain text for promptui to handle coloring
func FormatIssueDisplayWithColor(issue models.LinearIssue, teamKeys []string) string {
	// Just return the regular formatted display - promptui will handle colors
	return FormatIssueDisplay(issue)
}