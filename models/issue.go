package models

// Team represents a Linear team
type Team struct {
	Name        string `json:"name"`
	Key         string `json:"key"`
	DisplayName string `json:"displayName,omitempty"`
}

// State represents a Linear issue state
type State struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Cycle represents a Linear cycle
type Cycle struct {
	Name string `json:"name,omitempty"`
}

// User represents a Linear user
type User struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

// Project represents a Linear project
type Project struct {
	Name   string `json:"name"`
	Color  string `json:"color"`
	SlugId string `json:"slugId"`
	Id     string `json:"id"`
}

// LinearIssue represents a Linear issue
type LinearIssue struct {
	Id            string   `json:"id"`
	Title         string   `json:"title"`
	UpdatedAt     string   `json:"updatedAt"`
	Assignee      *User    `json:"assignee,omitempty"`
	Team          Team     `json:"team"`
	State         State    `json:"state"`
	Cycle         *Cycle   `json:"cycle,omitempty"`
	Description   string   `json:"description,omitempty"`
	BranchName    string   `json:"branchName"`
	CreatedAt     string   `json:"createdAt"`
	Estimate      *float64 `json:"estimate,omitempty"`
	Priority      *int     `json:"priority,omitempty"`
	PriorityLabel string   `json:"priorityLabel,omitempty"`
	StartedAt     string   `json:"startedAt,omitempty"`
	Creator       User     `json:"creator"`
	DueDate       string   `json:"dueDate,omitempty"`
	URL           string   `json:"url"`
	Project       *Project `json:"project,omitempty"`
}

// Define possible actions
type Action string

const (
	CopyBranchName Action = "copy-branch-name"
	OpenInBrowser  Action = "open-in-browser"
	CopyIssueURL   Action = "copy-issue-url"
)
