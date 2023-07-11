/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "database-populator",
	Short: "Populate any mongoDB database with the number of documents and structure",
	Long: `DatabasePopulator is a CLI library for Go that helps developers

populate any mongoDB with as many documents as they want.
This application is a tool to generate large quantities of documents for
testing tools on databases. The values of the generated documents can be specified by flags`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("host", "localhost", "The database host to connect to")
	generateCmd.Flags().String("port", "27017", "The database port to connect to")
	generateCmd.Flags().String("user", "root", "The database user to use for the connection")
	generateCmd.Flags().String("pass", "1234", "The database password to use for the connection")
	generateCmd.Flags().String("database", "", "The database name to use for the connection")
	generateCmd.Flags().String("path", "", "The schema path to use for documents")
	generateCmd.Flags().String("value", "", "The values to use for the document field \"<block_name.field_name> <value>...\" Needs to be space seperated and in quotes")
	generateCmd.MarkFlagsRequiredTogether("host", "port", "user", "pass", "database", "path")

	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().String("host", "localhost", "The database host to connect to")
	connectCmd.Flags().String("port", "27017", "The database port to connect to")
	connectCmd.Flags().String("user", "root", "The database user to use for the connection")
	connectCmd.Flags().String("pass", "1234", "The database password to use for the connection")
	connectCmd.MarkFlagsRequiredTogether("host", "port", "user", "pass")

}
