package linear

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zach/linear_cli_go/models"
)

const (
	LinearAPIURL = "https://api.linear.app/graphql"
)

type Client struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("LINEAR_API_KEY")
	}

	return &Client{
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type GraphQLResponse struct {
	Data   json.RawMessage `json:"data"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}

func (c *Client) ExecuteGraphQL(query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	req := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", LinearAPIURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", c.APIKey)

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var graphQLResp GraphQLResponse
	if err := json.Unmarshal(body, &graphQLResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if len(graphQLResp.Errors) > 0 {
		return nil, fmt.Errorf("GraphQL error: %s", graphQLResp.Errors[0].Message)
	}

	return &graphQLResp, nil
}

func (c *Client) GetIssues(onlyMine bool, projectID string, limit int) ([]models.LinearIssue, error) {
	var currentUserID string
	var err error
	
	if onlyMine {
		currentUserID, err = c.getCurrentUserID()
		if err != nil {
			return nil, fmt.Errorf("failed to get current user ID: %w", err)
		}
	}

	filterParts := []string{}
	if onlyMine && currentUserID != "" {
		filterParts = append(filterParts, fmt.Sprintf(`assignee: { id: { eq: "%s" } }`, currentUserID))
	}

	if projectID != "" {
		filterParts = append(filterParts, fmt.Sprintf(`project: { id: { eq: "%s" } }`, projectID))
	}

	filterArg := fmt.Sprintf("orderBy: updatedAt, first: %d", limit)
	if len(filterParts) > 0 {
		filterArg += fmt.Sprintf(`, filter: { %s }`, joinStrings(filterParts, ", "))
	}

	query := fmt.Sprintf(`
		query Issues { 
			issues(%s) {
				nodes { 
					id 
					title 
					description
					branchName
					createdAt
					updatedAt
					url
					team {
						name
						displayName
						key
					}
					state {
						name
						type
					}
					startedAt
					creator {
						name
						email
						displayName
					}
					dueDate
					cycle {
						name
					}
					dueDate
					estimate
					priority
					priorityLabel
					assignee { 
						name 
						displayName
						email
					} 
					project {
						name
						color
						slugId
						id
					}
				} 
			} 
		}
	`, filterArg)

	resp, err := c.ExecuteGraphQL(query, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Issues struct {
			Nodes []models.LinearIssue `json:"nodes"`
		} `json:"issues"`
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal issues: %w", err)
	}

	return result.Issues.Nodes, nil
}


func (c *Client) getCurrentUserID() (string, error) {
	query := `
		query GetCurrentUser { 
			viewer {
				id
			}
		}
	`

	resp, err := c.ExecuteGraphQL(query, nil)
	if err != nil {
		return "", err
	}

	var result struct {
		Viewer struct {
			ID string `json:"id"`
		} `json:"viewer"`
	}

	if err := json.Unmarshal(resp.Data, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal current user: %w", err)
	}

	return result.Viewer.ID, nil
}

func joinStrings(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += sep + strs[i]
	}
	return result
}
