package mock

import "github.com/zsiegel92/linear-cli-go/models"

func GetMockIssues() []models.LinearIssue {
	alice := &models.User{
		Name:        "Alice Smith",
		Email:       "alice@company.com",
		DisplayName: "Alice",
	}

	bob := &models.User{
		Name:        "Bob Johnson",
		Email:       "bob@company.com",
		DisplayName: "Bob",
	}

	charlie := &models.User{
		Name:        "Charlie Williams",
		Email:       "charlie@company.com",
		DisplayName: "Charlie",
	}

	frontendTeam := models.Team{
		Name: "Frontend",
		Key:  "FE",
	}

	backendTeam := models.Team{
		Name: "Backend",
		Key:  "BE",
	}

	designTeam := models.Team{
		Name: "Design",
		Key:  "DS",
	}

	backlogState := models.State{
		Name: "Backlog",
		Type: "backlog",
	}

	todoState := models.State{
		Name: "To Do",
		Type: "unstarted",
	}

	inProgressState := models.State{
		Name: "In Progress",
		Type: "started",
	}

	inReviewState := models.State{
		Name: "In Review",
		Type: "review",
	}

	doneState := models.State{
		Name: "Done",
		Type: "completed",
	}

	currentCycle := &models.Cycle{
		Name: "Sprint 24",
	}

	nextCycle := &models.Cycle{
		Name: "Sprint 25",
	}

	dashboardProject := &models.Project{
		Name:   "Dashboard Redesign",
		Color:  "#0000FF",
		SlugId: "dashboard-redesign",
		Id:     "proj_123",
	}

	authProject := &models.Project{
		Name:   "Auth System",
		Color:  "#FF0000",
		SlugId: "auth-system",
		Id:     "proj_456",
	}

	analyticsProject := &models.Project{
		Name:   "Analytics Platform",
		Color:  "#00FF00",
		SlugId: "analytics-platform",
		Id:     "proj_789",
	}

	estimate3 := float64(3)
	estimate5 := float64(5)
	estimate2 := float64(2)
	estimate8 := float64(8)

	priority1 := 1
	priority2 := 2
	priority3 := 3

	return []models.LinearIssue{
		{
			Id:            "LIN-101",
			Title:         "Implement new login form",
			UpdatedAt:     "2023-05-15T14:30:00Z",
			Assignee:      alice,
			Team:          frontendTeam,
			State:         inProgressState,
			Cycle:         currentCycle,
			Description:   "Create a new login form with improved validation and error handling.",
			BranchName:    "feature/login-form",
			CreatedAt:     "2023-05-10T09:00:00Z",
			Estimate:      &estimate3,
			Priority:      &priority2,
			PriorityLabel: "High",
			StartedAt:     "2023-05-12T10:15:00Z",
			Creator:       *bob,
			DueDate:       "2023-05-25T23:59:59Z",
			URL:           "https://linear.app/company/issue/LIN-101",
			Project:       authProject,
		},
		{
			Id:            "LIN-102",
			Title:         "Fix API rate limiting",
			UpdatedAt:     "2023-05-14T16:45:00Z",
			Assignee:      bob,
			Team:          backendTeam,
			State:         inReviewState,
			Cycle:         currentCycle,
			Description:   "Resolve issues with API rate limiting that causes errors during high traffic.",
			BranchName:    "fix/api-rate-limiting",
			CreatedAt:     "2023-05-09T13:20:00Z",
			Estimate:      &estimate5,
			Priority:      &priority1,
			PriorityLabel: "Urgent",
			StartedAt:     "2023-05-10T08:30:00Z",
			Creator:       *charlie,
			DueDate:       "2023-05-18T23:59:59Z",
			URL:           "https://linear.app/company/issue/LIN-102",
			Project:       authProject,
		},
		{
			Id:            "LIN-103",
			Title:         "Design dashboard widgets",
			UpdatedAt:     "2023-05-13T11:20:00Z",
			Assignee:      charlie,
			Team:          designTeam,
			State:         doneState,
			Cycle:         currentCycle,
			Description:   "Create design mockups for new dashboard widgets.",
			BranchName:    "design/dashboard-widgets",
			CreatedAt:     "2023-05-05T10:00:00Z",
			Estimate:      &estimate2,
			Priority:      &priority3,
			PriorityLabel: "Medium",
			StartedAt:     "2023-05-06T09:00:00Z",
			Creator:       *alice,
			DueDate:       "2023-05-12T23:59:59Z",
			URL:           "https://linear.app/company/issue/LIN-103",
			Project:       dashboardProject,
		},
		{
			Id:            "LIN-104",
			Title:         "Implement data visualization components",
			UpdatedAt:     "2023-05-16T15:10:00Z",
			Assignee:      alice,
			Team:          frontendTeam,
			State:         todoState,
			Cycle:         nextCycle,
			Description:   "Build reusable chart components for analytics dashboard.",
			BranchName:    "feature/data-viz",
			CreatedAt:     "2023-05-15T14:00:00Z",
			Estimate:      &estimate8,
			Priority:      &priority2,
			PriorityLabel: "High",
			StartedAt:     "",
			Creator:       *bob,
			DueDate:       "2023-06-01T23:59:59Z",
			URL:           "https://linear.app/company/issue/LIN-104",
			Project:       analyticsProject,
		},
		{
			Id:            "LIN-105",
			Title:         "Setup database migrations",
			UpdatedAt:     "2023-05-12T08:30:00Z",
			Assignee:      nil,
			Team:          backendTeam,
			State:         backlogState,
			Cycle:         nil,
			Description:   "Configure automated database migration process.",
			BranchName:    "chore/db-migrations",
			CreatedAt:     "2023-05-01T11:45:00Z",
			Estimate:      nil,
			Priority:      nil,
			PriorityLabel: "",
			StartedAt:     "",
			Creator:       *bob,
			DueDate:       "",
			URL:           "https://linear.app/company/issue/LIN-105",
			Project:       nil,
		},
	}
}

func GetMockIssuesForUser(userDisplayName string) []models.LinearIssue {
	allIssues := GetMockIssues()
	var filteredIssues []models.LinearIssue
	
	for _, issue := range allIssues {
		if issue.Assignee != nil && issue.Assignee.DisplayName == userDisplayName {
			filteredIssues = append(filteredIssues, issue)
		}
	}
	
	return filteredIssues
}
