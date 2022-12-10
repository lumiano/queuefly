package infra

import (
	"go.uber.org/fx"
)

var Modules = fx.Options(fx.Provide(NewEchoHandler), fx.Provide(NewRequestHandler), fx.Provide(NewConfig))
