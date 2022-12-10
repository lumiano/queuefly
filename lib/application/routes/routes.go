package routes

import "go.uber.org/fx"

var Modules = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(UserRoutes UserRoutes) Routes {
	return Routes{UserRoutes}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
