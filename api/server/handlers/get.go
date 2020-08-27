package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//HandleGETAllIssues - Route to retrieve all issues
// @summary Retrieves all existing issues
// @description Retrieves all issues
// @tags Retrieval
// @accept json
// @produce json
// @success 200 {object} models.IssueResponse
// @failure 400 {object} models.ErrorWrapper
// @failure 404 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issues [get]
func HandleGETAllIssues(storage persistence.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[GET] get-all-issues")

		issuesResponse, err := storage.RetrieveIssues()

		if err == sql.ErrNoRows {
			models.SetErrorStatusJSON(c, http.StatusNotFound, "could not find issue")
			return
		}

		if err != nil {
			l.Errorf("error retrieving issue in db: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		l.Debug("issues successfully retrieved")
		c.JSON(200, issuesResponse)
		return
	}
}

//HandleGETByID - Route to retrieve an issue by the issue id
// @summary Retrieves an issue given issue id
// @description Retrieves an issue given issue id
// @tags Retrieval
// @accept json
// @produce json
// @Param id path int true "ID of the issue"
// @success 200 {object} models.IssueResponse
// @failure 400 {object} models.ErrorWrapper
// @failure 404 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issue/{id} [get]
func HandleGETByID(storage persistence.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[GET] get-issue-by-id")

		issueID, err := strconv.ParseInt(c.Param("issueID"), 10, 64)
		if err != nil {
			models.SetErrorStatusJSON(c, http.StatusBadRequest, "invalid issue id format")
			return
		}

		issueResponse, err := storage.RetrieveIssueByID(issueID)

		if err == sql.ErrNoRows {
			models.SetErrorStatusJSON(c, http.StatusNotFound, "could not find issue")
			return
		}

		if err != nil {
			l.Errorf("error retrieving issue in db: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		l.Debug("issue successfully retrieved")
		c.JSON(200, issueResponse)
		return
	}
}

//HandleGETByStatus - Route to retrieve an issue filtered by the status
// @summary Retrieves an issue given status
// @description Retrieves an issue given status (open, closed, in progress)
// @tags Retrieval
// @accept json
// @produce json
// @param status query models.StatusQueryParam false "issue priority request"
// @success 200 {object} models.IssueResponse
// @failure 400 {object} models.ErrorWrapper
// @failure 404 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issues/status [get]
func HandleGETByStatus(storage persistence.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[GET] get-issue-by-status")

		var statusQuery models.StatusQueryParam
		err := c.ShouldBindQuery(&statusQuery)
		if err != nil {
			models.SetErrorStatusJSON(c, http.StatusBadRequest, "could not filter by status")
			return
		}

		issueResponse, err := storage.RetrieveIssueByStatus(statusQuery.Status)

		if err == sql.ErrNoRows {
			models.SetErrorStatusJSON(c, http.StatusNotFound, "could not find issue")
			return
		}

		if err != nil {
			l.Errorf("error retrieving issue in db: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		l.Debug("issue successfully retrieved")
		c.JSON(200, issueResponse)
		return
	}
}

//HandleGETByPriority - Route to retrieve an issue filtered by the priority
// @summary Retrieves an issue given priority
// @description Retrieves an issue given priority
// @tags Retrieval
// @accept json
// @produce json
// @param start query models.PriorityQueryParam false "priority start bound"
// @param end query models.PriorityQueryParam false "priority end bound"
// @success 200 {object} models.IssueResponse
// @failure 400 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issues/priority [get]
func HandleGETByPriority(storage persistence.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[GET] get-issue-by-priority")

		var priorityQuery models.PriorityQueryParam
		err := c.ShouldBindQuery(&priorityQuery)
		if err != nil {
			models.SetErrorStatusJSON(c, http.StatusBadRequest, "could not filter by priority")
			return
		}

		issueResponse, err := storage.RetrieveIssueByPriority(priorityQuery.PriorityStart, priorityQuery.PriorityEnd)

		if err != nil {
			l.Errorf("error retrieving issue in db: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		l.Debug("issue successfully retrieved")
		c.JSON(200, issueResponse)
		return
	}
}
