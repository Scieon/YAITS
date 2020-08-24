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

func HandleGETByID(storage persistence.MysqlStorage) gin.HandlerFunc {
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

func HandleGETByStatus(storage persistence.MysqlStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[GET] get-issue-by-status")

		var statusQuery models.StatusQueryParam
		err := c.ShouldBindQuery(&statusQuery)
		if err != nil {
			models.SetErrorStatusJSON(c, http.StatusBadRequest, "could not filter by status")
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
