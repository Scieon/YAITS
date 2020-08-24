package persistence

import "github.com/YAITS/api/models"

type Storage struct{}

const (
	IssueID     = int64(1)
	Summary     = "This is a summary"
	Description = "This is a description"
	Assignee    = "John Doe"
	Status      = "Open"
	Priority    = int64(1)
	CreateDate  = "some date"
)

var MockIssueResponse = models.IssueResponse{
	ID:          IssueID,
	Description: Description,
	Summary:     Summary,
	Assignee:    Assignee,
	CreateDate:  CreateDate,
	Priority:    Priority,
	Status:      Status,
}

func (storage *Storage) CreateIssue(_, _, _ string, _ int64) (int64, error) {
	return 1, nil
}

func (storage *Storage) UpdateIssue(_, _, _, _ string, _, _ int64) (*models.IssueResponse, error) {
	return &MockIssueResponse, nil
}

func (storage *Storage) RetrieveIssueByID(_ int64) (models.IssueResponse, error) {
	return MockIssueResponse, nil
}

func (storage *Storage) RetrieveIssueByStatus(_ string) ([]models.IssueResponse, error) {
	return []models.IssueResponse{MockIssueResponse}, nil
}

func (storage *Storage) RetrieveIssueByPriority(_, _ int64) ([]models.IssueResponse, error) {
	return []models.IssueResponse{MockIssueResponse}, nil
}

func (storage *Storage) DeleteIssueByID(_ int64) error {
	return nil
}

func NewMockStorage() *Storage {
	return &Storage{}
}
