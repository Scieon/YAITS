package persistence

import (
	"database/sql"
	"github.com/YAITS/api/models"
)

//MysqlStorage - Hold sql database pointer
type MysqlStorage struct {
	db *sql.DB
}

//NewMysqlStorage - Create MysqlStorage object
func NewMysqlStorage(db *sql.DB) *MysqlStorage {
	return &MysqlStorage{db: db}
}

// IssueEntry is a struct containing all issue information
type IssueEntry struct {
	ID          int
	Description string
	Summary     string
	Status      string
	Assignee    string
	CreateDate  string
	Priority    int
}

// CreateIssue creates a new issue
func (mysqlSt *MysqlStorage) CreateIssue(summary, description, assignee string, priority int64) (int64, error) {
	var err error
	var result sql.Result
	if assignee == "" {
		insertQuery := "INSERT INTO issues(summary, description, priority) VALUES(?, ?, ?)"
		result, err = mysqlSt.db.Exec(insertQuery, summary, description, priority)
	} else {
		insertQuery := "INSERT INTO issues(summary, description, priority, assignee) VALUES(?, ?, ?, ?)"
		result, err = mysqlSt.db.Exec(insertQuery, summary, description, priority, assignee)
	}

	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()
	return id, nil
}

// UpdateIssue edits an existing issue
func (mysqlSt *MysqlStorage) UpdateIssue(summary, description, assignee, status string, priority, issueID int64) (*models.IssueResponse, error) {
	var err error

	issue, err := mysqlSt.RetrieveIssueByID(issueID)
	if err != nil {
		return nil, err
	}

	if summary != "" {
		issue.Summary = summary
	}
	if description != "" {
		issue.Description = description
	}
	if assignee != "" {
		issue.Assignee = assignee
	}
	if status != "" {
		issue.Status = status
	}
	if priority != 0 {
		issue.Priority = priority
	}

	updateQuery := "UPDATE issues SET summary = ?, description = ?, assignee = ?, status = ?, priority = ? WHERE id = ?"

	_, err = mysqlSt.db.Exec(updateQuery, issue.Summary, issue.Description, issue.Assignee, issue.Status, issue.Priority, issueID)
	if err != nil {
		return nil, err
	}

	return &issue, nil
}

// RetrieveIssueByID returns an issue filtered by the issue id
func (mysqlSt *MysqlStorage) RetrieveIssueByID(issueID int64) (models.IssueResponse, error) {
	var resp models.IssueResponse
	var summary, description, status, assignee, createDate string
	var id, priority int

	query := `SELECT id, summary, description, priority, status, assignee, createDate FROM issues WHERE id = ?`

	err := mysqlSt.db.QueryRow(query, issueID).Scan(&id, &summary, &description, &priority, &status, &assignee, &createDate)

	if err != nil {
		return resp, err
	}

	resp = models.IssueResponse{
		ID:          int64(id),
		Summary:     summary,
		Description: description,
		Priority:    int64(priority),
		Status:      status,
		Assignee:    assignee,
		CreateDate:  createDate,
	}

	return resp, nil
}

// RetreiveIssueByStatus returns an issue filtered by the status (open, closed, in progress)
func (mysqlSt *MysqlStorage) RetrieveIssueByStatus(statusFilter string) ([]models.IssueResponse, error) {
	resp := make([]models.IssueResponse, 0)
	var summary, status, description, assignee, createDate string
	var id, priority int

	query := `SELECT id, summary, description, priority, status, assignee, createDate FROM issues WHERE status = ?`

	rows, err := mysqlSt.db.Query(query, statusFilter)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&id, &summary, &description, &priority, &status, &assignee, &createDate)
		if err != nil {
			return nil, err
		}

		resp = append(resp, models.IssueResponse{
			ID:          int64(id),
			Summary:     summary,
			Description: description,
			Priority:    int64(priority),
			Status:      status,
			Assignee:    assignee,
			CreateDate:  createDate,
		})
	}

	return resp, nil
}

// DeleteIssueByID deletes an issue filtered by the issue id
func (mysqlSt *MysqlStorage) DeleteIssueByID(issueID int64)  error {

	query := `DELETE FROM issues WHERE id = ?`

	_, err := mysqlSt.db.Exec(query, issueID)

	if err != nil {
		return err
	}

	return nil
}
