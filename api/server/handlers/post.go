package handlers

import (
	"github.com/YAITS/api/models"
	"github.com/YAITS/api/persistence"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

		err = storage.CreateIssue(req.Summary, req.Description, "", req.Priority)

		if err != nil {
			l.Errorf("couldn't insert into db: %s", err.Error())
			c.JSON(http.StatusInternalServerError, "db error")
			return
		}

		l.Debug("insertion successful")
		c.JSON(http.StatusCreated, "ok")
		return
	}
}
