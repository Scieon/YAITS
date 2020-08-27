package handlers

import (
	"net/http"
	"strconv"

	"go.uber.org/zap"

	_ "github.com/YAITS/api/docs"
	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
)

//HandleDELETE - Route to delete an issue
// @summary Delete an issue
// @description Deletes an issue given an issue id
// @tags Deletion
// @accept json
// @produce json
// @Param id path int true "ID of the issue"
// @success 204 {} No Content
// @failure 400 {object} models.ErrorWrapper
// @failure 404 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issue/{id} [delete]
func HandleDELETE(storage persistence.Storage) gin.HandlerFunc {
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
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		l.Debug("issue deleted")
		c.JSON(http.StatusNoContent,  "issue deleted")
		return
	}
}
