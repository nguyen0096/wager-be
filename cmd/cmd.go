package cmd

import "github.com/spf13/cobra"

func New() *cobra.Command {
	root := &cobra.Command{
		Use: "wager",
	}

	root.AddCommand(&cobra.Command{
		Use:   "server",
		Short: "start wager server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	})

	return root
}
