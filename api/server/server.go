package server

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/YAITS/api/persistence"
	"github.com/YAITS/api/server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewServer(address string, logger *zap.SugaredLogger, storage *persistence.MysqlStorage) *http.Server {
	router := BuildRouter(logger, *storage)
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}

func BuildRouter(logger *zap.SugaredLogger, storage persistence.MysqlStorage) *gin.Engine {
	router := gin.New()

	router.Use(setupLogger(logger))
	router.Use(gin.Recovery())

	apiGroup := router.Group("/api")

	apiGroup.GET("/issue/:issueID", handlers.HandleGETByID(storage))
	apiGroup.GET("/issues/status", handlers.HandleGETByStatus(storage))
	apiGroup.GET("/issues/priority", handlers.HandleGETByPriority(storage))

	apiGroup.POST("/issue", handlers.HandlePOST(storage))

	apiGroup.PATCH("/issue/:issueID", handlers.HandlePATCH(storage))

	apiGroup.DELETE("/issue/:issueID", handlers.HandleDELETE(storage))

	return router
}

func setupLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(context *gin.Context) {
		reqID := uuid.New().String()
		loggerWithReqID := logger.With("request-id", reqID)
		context.Set("logger", loggerWithReqID)
	}
}
