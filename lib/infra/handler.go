package infra

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler(logger *EchoHandler) RequestHandler {

	engine := gin.New()

	// Add a gin zap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	//zap, _ := logger.
	engine.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	engine.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	return RequestHandler{Gin: engine}
}
