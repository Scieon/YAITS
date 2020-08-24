package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
)

//HandlePATCH - Route to update an issue
// @summary Update an issue
// @description Updates an issue given an issue id
// @tags issue
// @accept json
// @produce json
// @Param id path int true "ID of the issue"
// @param updateIssueRequest body models.UpdateIssueRequest true "YAITS update request"
// @success 200 {object} models.IssueResponse
// @failure 400 {object} models.ErrorWrapper
// @failure 404 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issue/{id} [patch]
func HandlePATCH(storage persistence.MysqlStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[PATCH] update-issue")

		// Retrieve issue to update
		issueID, err := strconv.ParseInt(c.Param("issueID"), 10, 64)
		if err != nil {
			models.SetErrorStatusJSON(c, http.StatusBadRequest, "invalid issue id format")
			return
		}

		var req models.UpdateIssueRequest
		err = c.ShouldBindJSON(&req)

		l = l.With("request", req, "issueID", issueID)
		l.Debug("received issue update request")

		if err != nil {
			l.Errorf("couldn't bind to issue request: %s", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		issue, err := storage.UpdateIssue(req.Summary, req.Description, req.Assignee, req.Status, req.Priority, issueID)

		if err == sql.ErrNoRows {
			models.SetErrorStatusJSON(c, http.StatusNotFound, "could not find issue")
			return
		}

		if err != nil {
			l.Errorf("couldn't update: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, "db error")
			return
		}

		l.Debug("update successful")
		c.JSON(http.StatusOK, issue)
		return
	}
}
