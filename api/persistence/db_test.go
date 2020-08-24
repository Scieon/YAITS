package persistence

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/require"
)

const (
	IssueID     = int64(1)
	Summary     = "This is a summary"
	Description = "This is a description"
	Assignee    = "John Doe"
	Status      = "Open"
	Priority    = int64(1)
	CreateDate  = "some date"
)

func TestMysqlStorage_RetrieveIssueByID(t *testing.T) {
	// setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// force db closure at end of test
	defer func() {
		_ = db.Close()
	}()

	testingStorage := NewMysqlStorage(db)

	mock.ExpectQuery("SELECT (.+) FROM issues").
		WithArgs(IssueID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "summary", "description", "priority", "status", "assignee", "createDate"}).
			AddRow(IssueID, Summary, Description, Priority, Status, Assignee, CreateDate))

	// run the code
	if _, err = testingStorage.RetrieveIssueByID(IssueID); err != nil {
		t.Errorf("Error should not have occurred while updating issue: %s", err)
	}

	//check expectations are met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %s", err)
	}
}

func TestMysqlStorage_RetrieveIssueByStatus(t *testing.T) {
	// setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// force db closure at end of test
	defer func() {
		_ = db.Close()
	}()

	testingStorage := NewMysqlStorage(db)

	mock.ExpectQuery("SELECT (.+) FROM issues WHERE status").
		WithArgs(Status).
		WillReturnRows(sqlmock.NewRows([]string{"id", "summary", "description", "priority", "status", "assignee", "createDate"}).
			AddRow(IssueID, Summary, Description, Priority, Status, Assignee, CreateDate))

	// run the code
	if _, err = testingStorage.RetrieveIssueByStatus(Status); err != nil {
		t.Errorf("Error should not have occurred while updating issue: %s", err)
	}

	//check expectations are met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations not met: %s", err)
	}
}

func TestMysqlStorage_UpdateIssue(t *testing.T) {
	// setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// force db closure at end of test
	defer func() {
		_ = db.Close()
	}()

	testingStorage := NewMysqlStorage(db)

	t.Run("NoError", func(t *testing.T) {
		mock.ExpectQuery("SELECT (.+) FROM issues").
			WithArgs(IssueID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "summary", "description", "priority", "status", "assignee", "createDate"}).
				AddRow(IssueID, Summary, Description, Priority, Status, Assignee, CreateDate))

		// set expectations
		mock.ExpectExec("UPDATE issues SET").
			WithArgs(Summary, Description, Assignee, Status, Priority, IssueID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// run the code
		if _, err = testingStorage.UpdateIssue(Summary, Description, Assignee, Status, Priority, IssueID); err != nil {
			t.Errorf("Error should not have occurred while updating issue: %s", err)
		}

		//check expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})
}

func TestMysqlStorage_DeleteIssueByID(t *testing.T) {
	// setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	// force db closure at end of test
	defer func() {
		_ = db.Close()
	}()

	testingStorage := NewMysqlStorage(db)

	t.Run("NoError", func(t *testing.T) {
		// set expectations
		mock.ExpectExec("DELETE FROM issues").
			WithArgs(IssueID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// run the code
		if err = testingStorage.DeleteIssueByID(IssueID); err != nil {
			t.Errorf("Error should not have occurred while deleting issue: %s", err)
		}

		//check expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		// set expectations
		mock.ExpectExec("DELETE FROM issues").
			WithArgs(IssueID).
			WillReturnError(errors.New("err"))

		// run the code
		if err = testingStorage.DeleteIssueByID(IssueID); err == nil {
			t.Errorf("Error should have occured while deleting issue: %s", err)
		}

		//check expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})
}

func TestMysqlStorage_RetrieveIssueByPriority(t *testing.T) {
	// setup
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	priorityStart := int64(1)
	priorityEnd := int64(2)

	// force db closure at end of test
	defer func() {
		_ = db.Close()
	}()

	testingStorage := NewMysqlStorage(db)

	t.Run("NoErrorOnlyPriorityStart", func(t *testing.T) {
		// set expectations
		mock.ExpectQuery("SELECT (.+) FROM issues").
			WithArgs(priorityStart).
			WillReturnRows(sqlmock.NewRows([]string{"id", "summary", "description", "priority", "status", "assignee", "createDate"}).
				AddRow(IssueID, Summary, Description, Priority, Status, Assignee, CreateDate))

		// run the code
		if _, err = testingStorage.RetrieveIssueByPriority(priorityStart, 0); err != nil {
			t.Errorf("Error should not have occurred while retrieving issue: %s", err)
		}

		//check expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})

	t.Run("NoErrorPriorityStartAndEnd", func(t *testing.T) {
		// set expectations
		mock.ExpectQuery("SELECT (.+) FROM issues").
			WithArgs(priorityStart, priorityEnd).
			WillReturnRows(sqlmock.NewRows([]string{"id", "summary", "description", "priority", "status", "assignee", "createDate"}).
				AddRow(IssueID, Summary, Description, Priority, Status, Assignee, CreateDate))

		// run the code
		if _, err = testingStorage.RetrieveIssueByPriority(priorityStart, priorityEnd); err != nil {
			t.Errorf("Error should not have occurred while retrieving issue: %s", err)
		}

		//check expectations are met
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Expectations not met: %s", err)
		}
	})
}
