package cmd

import "github.com/spf13/cobra"

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a database",
	Long:  ``,
	Run:   connect,
}

func connect(cmd *cobra.Command, args []string) {
}
