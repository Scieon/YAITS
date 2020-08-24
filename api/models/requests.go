package models

// NewIssueRequest is the incoming request to create a new issue
type NewIssueRequest struct {
	Description string `json:"description" binding:"required"`
	Summary     string `json:"summary" binding:"required"`
	Priority    int64  `json:"priority" binding:"required"`
	Assignee    string `json:"assignee"`
}

// UpdateIssueRequest is the incoming request to update an existing issue
type UpdateIssueRequest struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Priority    int64  `json:"priority"`
	Assignee    string `json:"assignee"`
	Status      string `json:"status"`
}

// StatusQueryParam is the query header parameter to filter issues by statuses
type StatusQueryParam struct {
	Status string `json:"status"`
}
