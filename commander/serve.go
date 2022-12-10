package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"queuefly/lib/application/middlewares"
	"queuefly/lib/application/routes"
	"queuefly/lib/data"
	"queuefly/lib/infra"
)

// ServeCommand test command
type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() CommandRunner {
	return func(
		middlewares middlewares.Middlewares,
		environment infra.Config,
		handler infra.RequestHandler,
		routes routes.Routes,
		logger *infra.EchoHandler,
		database data.Database,

	) {
		middlewares.Setup()
		routes.Setup()

		zap.ReplaceGlobals(logger.Logger)

		logger.Info("Running server")
		_ = handler.Gin.Run(fmt.Sprintf(":%s", environment.ServerPort))
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
