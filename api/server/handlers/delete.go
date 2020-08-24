package handlers

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
)

func HandleDELETE(storage persistence.MysqlStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[DELETE] delete-issue")

		// Retrieve issue to delete
		issueID, err := strconv.ParseInt(c.Param("issueID"), 10, 64)
		if err != nil {
			models.SetErrorStatusJSON(c, http.StatusBadRequest, "invalid issue id format")
			return
		}

		l = l.With( "issueID", issueID)
		l.Debug("received issue deletion request")

		err = storage.DeleteIssueByID(issueID)

		if err != nil {
			l.Errorf("couldn't update: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, "db error")
			return
		}

		l.Debug("issue deleted")
		c.JSON(http.StatusNoContent,  "issue deleted")
		return
	}
}
