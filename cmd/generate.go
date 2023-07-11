package cmd

import "github.com/spf13/cobra"

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates documents for a MongoDB database",
	Long:  ``,
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	//TODO: Generate documents here
}
