package commands

import (
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"queuefly/lib/infra"
)

var cmd = map[string]Command{
	"app:serve": NewServeCommand(),
}

// GetSubCommands gives a list of sub commands
func GetSubCommands(opt fx.Option) []*cobra.Command {
	var subCommands []*cobra.Command
	for name, cmd := range cmd {
		subCommands = append(subCommands, WrapSubCommand(name, cmd, opt))
	}
	return subCommands
}

type FxLogger struct {
	*zap.Logger
}

func WrapSubCommand(name string, cmd Command, opt fx.Option) *cobra.Command {
	wrappedCmd := &cobra.Command{
		Use:   name,
		Short: cmd.Short(),
		Run: func(c *cobra.Command, args []string) {
			logger := infra.NewEchoHandler()

			opts := fx.Options(
				fx.WithLogger(func() fxevent.Logger {
					return &fxevent.ZapLogger{Logger: logger.Logger}

				}),
				fx.Invoke(cmd.Run()),
			)
			ctx := context.Background()
			app := fx.New(opt, opts)
			err := app.Start(ctx)
			defer app.Stop(ctx)
			if err != nil {
				logger.Fatal(err.Error())
			}
		},
	}
	cmd.Setup(wrappedCmd)
	return wrappedCmd
}
