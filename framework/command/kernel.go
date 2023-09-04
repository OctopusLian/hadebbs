package command

import "hadebbs/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	root.AddCommand(DemoCommand)

	root.AddCommand(initAppCommand())
}
