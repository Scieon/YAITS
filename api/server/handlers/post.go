package handlers

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
)

func HandlePOST(storage persistence.MysqlStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := c.MustGet("logger").(*zap.SugaredLogger).With("handler", "[POST] create-issue")

		var req models.NewIssueRequest
		err := c.ShouldBindJSON(&req)

		l = l.With("request", req)
		l.Debug("received issue creation request")

		if err != nil {
			l.Errorf("couldn't bind to issue request: %s", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		id, err := storage.CreateIssue(req.Summary, req.Description, req.Assignee, req.Priority)

		if err != nil {
			l.Errorf("couldn't insert into db: %s", err.Error())
			c.JSON(http.StatusInternalServerError, "db error")
			return
		}

		l.Debug("insertion successful")
		c.JSON(http.StatusCreated, id)
		return
	}
}
