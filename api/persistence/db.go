package persistence

import (
	"database/sql"
)


//MysqlStorage - Hold sql database pointer
type MysqlStorage struct {
	db *sql.DB
}

//NewMysqlStorage - Create MysqlStorage object
func NewMysqlStorage(db *sql.DB) *MysqlStorage {
	return &MysqlStorage{db: db}
}

// todo missing reporter
func (mysqlSt *MysqlStorage) CreateIssue(summary, description, assignee string, priority int) error {
	if assignee == "" {
		insertQuery := "INSERT INTO issues(summary, description, priority) VALUES(?, ?, ?, ?)"

		_, err := mysqlSt.db.Exec(insertQuery, summary, description, priority)

		if err != nil {
			return err
		}
	}

	return nil
}
