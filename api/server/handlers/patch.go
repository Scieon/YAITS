package handlers

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
)

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

		if err != nil {
			l.Errorf("couldn't update: %s", err.Error())
			c.JSON(http.StatusInternalServerError, "db error")
			return
		}

		l.Debug("update successful")
		c.JSON(http.StatusCreated, issue)
		return
	}
}
