package command

import "github.com/wxc/webfmk/framework/cobra"

func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)
	root.AddCommand(initAppCommand())
}
