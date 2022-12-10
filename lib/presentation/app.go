package presentation

import (
	"github.com/spf13/cobra"
	commands "queuefly/commander"
)

var rootCmd = &cobra.Command{
	Use:   "clean-gin",
	Short: "Queuefly Application",
	Long: `
█▀▀ █░░ █▀▀ ▄▀█ █▄░█ ▄▄ █▀▀ █ █▄░█
█▄▄ █▄▄ ██▄ █▀█ █░▀█ ░░ █▄█ █ █░▀█      
                                         		
This is a command runner or cli for api architecture in golang. 
Using this we can use underlying dependency injection container for running scripts. 
Main advantage is that, we can use same services, repositories, infrastructure present in the application itself`,
	TraverseChildren: true,
}

// App root of application
type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}
	cmd.AddCommand(commands.GetSubCommands(AppModules)...)
	return cmd
}

var RootApp = NewApp()
