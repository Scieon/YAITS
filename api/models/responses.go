package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IssueResponse contains all information about an issue
type IssueResponse struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
	CreateDate  string `json:"createDate"`
	Priority    int64  `json:"priority"`
}

// ErrorWrapper provides a general template for the response
type ErrorWrapper struct {
	Errors []StandardError `json:"errors"`
}

// StandardError is the normal error to be returned
type StandardError struct {
	Code        int    `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

// NewErrorWrapper returns an ErrorWrapper with the appropriate parameters
func NewErrorWrapper(code int, description string) ErrorWrapper {
	return ErrorWrapper{Errors: []StandardError{{Code: code, Title: http.StatusText(code), Description: description}}}
}

// SetErrorStatusJSON generates and sends an error with a description
func SetErrorStatusJSON(c *gin.Context, status int, description string) {
	c.JSON(status, NewErrorWrapper(status, description))
}
