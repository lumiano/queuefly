package presentation

import (
	"go.uber.org/fx"
	"queuefly/lib/application/controllers"
	"queuefly/lib/application/middlewares"
	"queuefly/lib/application/routes"
	"queuefly/lib/data"
	"queuefly/lib/domain/repositories"
	"queuefly/lib/domain/services"
	"queuefly/lib/infra"
)

var AppModules = fx.Options(controllers.Modules, routes.Modules, infra.Modules, data.Modules, services.Modules, middlewares.Modules, repositories.Modules)
