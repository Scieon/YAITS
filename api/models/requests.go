package models

type NewIssueRequest struct {
	Description string `form:"description" binding:"required"`
	Summary     string `form:"summary" binding:"required"`
	Priority    int    `form:"priority" binding:"required"`
	Assignee    string `form:"assignee"`
}
