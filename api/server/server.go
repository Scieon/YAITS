package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

func NewServer(address string, logger *zap.SugaredLogger) *http.Server {
	router := BuildRouter(logger)
	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}

func BuildRouter(logger *zap.SugaredLogger) *gin.Engine {
	router := gin.New()

	router.Use(setupLogger(logger))
	router.Use(gin.Recovery())

	apiGroup := router.Group("/api")

	apiGroup.GET("/", func(c *gin.Context) {
		logger.Debug("[GET] Test handler")
		c.String(http.StatusOK, "hello world")
	})

	return router
}

func setupLogger(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(context *gin.Context) {
		reqID := uuid.New().String()
		loggerWithReqID := logger.With("request-id", reqID)
		context.Set("logger", loggerWithReqID)
	}
}
