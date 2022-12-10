package routes

import (
	"queuefly/lib/application/controllers"
	"queuefly/lib/infra"
)

type UserRoutes struct {
	logger  *infra.EchoHandler
	handler infra.RequestHandler
	controllers.UserController
}

func (s UserRoutes) Setup() {
	s.logger.Info("Setup Users Routes")

	api := s.handler.Gin.Group("/api")
	{
		api.POST("/user", s.UserController.CreateUser)
	}

}

func NewUserRoutes(Logger *infra.EchoHandler, Handler infra.RequestHandler, UserController controllers.UserController) UserRoutes {
	return UserRoutes{Logger, Handler, UserController}
}
