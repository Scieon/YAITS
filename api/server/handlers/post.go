package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
)

//HandlePOST - Route to create an issue
// @summary Create an issue
// @description Create a new issue
// @tags Creation
// @accept json
// @produce json
// @param issueRequest body models.NewIssueRequest true "YAITS creation request"
// @success 201 {object} models.IssueIDResponse
// @failure 400 {object} models.ErrorWrapper
// @failure 500 {object} models.ErrorWrapper
// @router /issue [post]
func HandlePOST(storage persistence.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[POST] create-issue")

		var req models.NewIssueRequest
		err := c.ShouldBindJSON(&req)

		l = l.With("request", req)
		l.Debug("received issue creation request")

		if err != nil {
			l.Errorf("couldn't bind to issue request: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		id, err := storage.CreateIssue(req.Summary, req.Description, req.Assignee, req.Priority)

		if err != nil {
			l.Errorf("couldn't insert into db: %s", err.Error())
			models.SetErrorStatusJSON(c, http.StatusInternalServerError, err.Error())
			return
		}

		l.Debug("insertion successful")
		c.JSON(http.StatusCreated, models.IssueIDResponse{ID: id})
		return
	}
}
