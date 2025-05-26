package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/zsiegel92/linear-cli-go/models"
)

func GetSlug(text string) string {
	preserved := ":()[]{}-&"
	preservedSet := make(map[rune]bool)
	for _, char := range preserved {
		preservedSet[char] = true
	}

	if text == "" {
		return ""
	}

	text = strings.TrimSpace(text)
	words := strings.Fields(text)
	
	var result strings.Builder
	
	for _, word := range words {
		if word == "" {
			continue
		}
		
		haveNonPreservedChar := false
		var accepted []rune
		
		for _, char := range word {
			if preservedSet[char] {
				accepted = append(accepted, char)
			} else if !haveNonPreservedChar {
				haveNonPreservedChar = true
				accepted = append(accepted, char)
			}
		}
		
		result.WriteString(string(accepted))
	}
	
	return result.String()
}


func FormatTimeAgo(timeStr string) string {
	if timeStr == "" {
		return ""
	}

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}

	duration := time.Since(t)
	days := int(duration.Hours() / 24)

	if days == 0 {
		hours := int(duration.Hours())
		if hours == 0 {
			minutes := int(duration.Minutes())
			if minutes == 0 {
				return "just now"
			}
			return fmt.Sprintf("%d minutes ago", minutes)
		}
		return fmt.Sprintf("%d hours ago", hours)
	} else if days == 1 {
		return "1 day ago"
	} else if days < 30 {
		return fmt.Sprintf("%d days ago", days)
	} else if days < 365 {
		months := days / 30
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	} else {
		years := days / 365
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}
}

func FormatEstimate(estimate *float64) string {
	if estimate == nil {
		return ""
	}
	if *estimate == float64(int(*estimate)) {
		return fmt.Sprintf("%.0f", *estimate)
	}
	return fmt.Sprintf("%.1f", *estimate)
}

func FormatIssueDisplay(issue models.LinearIssue) string {
	assignee := "UNASSIGNED"
	if issue.Assignee != nil {
		assignee = issue.Assignee.DisplayName
	}

	team := issue.Team.Key

	projectSlug := ""
	if issue.Project != nil {
		projectSlug = strings.ToUpper(GetSlug(issue.Project.Name))
	}

	prefix := fmt.Sprintf("[%s - %s", assignee, team)
	if projectSlug != "" {
		prefix += " - " + projectSlug
	}
	prefix += "]"

	estimate := ""
	if issue.Estimate != nil {
		estimate = fmt.Sprintf(" (%s)", FormatEstimate(issue.Estimate))
	}

	timeAgo := ""
	if issue.UpdatedAt != "" {
		timeAgo = fmt.Sprintf(" (%s)", FormatTimeAgo(issue.UpdatedAt))
	}

	return fmt.Sprintf("%s%s %s%s", prefix, estimate, issue.Title, timeAgo)
}