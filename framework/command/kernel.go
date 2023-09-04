package command

import "hadebbs/framework/cobra"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	//root.AddCommand(DemoCommand)

	// env
	root.AddCommand(initEnvCommand())
	// app
	root.AddCommand(initAppCommand())
	// cron
	root.AddCommand(initCronCommand())
}
