package console

import (
	"github.com/wxc/webfmk/app/console/command/demo"
	"github.com/wxc/webfmk/framework"
	"github.com/wxc/webfmk/framework/cobra"
	"github.com/wxc/webfmk/framework/command"
)

func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "fmk",
		Short: "fmk command",
		Long:  "fmk framwork is a cli tool",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}

	rootCmd.SetContainer(container)
	command.AddKernelCommands(rootCmd)
	AddAppCommand(rootCmd)

	return rootCmd.Execute()
}

func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}
