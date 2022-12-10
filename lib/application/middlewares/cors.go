package middlewares

import (
	cors "github.com/gin-contrib/cors"
	"queuefly/lib/infra"
	"time"
)

type CorsMiddleware struct {
	handler infra.RequestHandler
	logger  *infra.EchoHandler
	config  infra.Config
}

func NewCorsMiddleware(handler infra.RequestHandler, logger *infra.EchoHandler, config infra.Config) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		logger:  logger,
		config:  config,
	}
}

// Setup sets up cors middleware
func (m CorsMiddleware) Setup() {
	m.logger.Info("Setting up cors middleware")

	m.handler.Gin.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowHeaders:     []string{"*"},
		MaxAge:           12 * time.Hour,
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	}))
}
