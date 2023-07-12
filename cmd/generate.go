package cmd

import (
	"fmt"
	"os"
	"strconv"

	database "github.com/polBachelin/database-populator/internal/database"
	"github.com/polBachelin/database-populator/internal/generation"
	"github.com/polBachelin/database-populator/internal/schema"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates documents for a MongoDB database",
	Long:  ``,
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Please provide the amount of documents you want to generate")
		return
	}
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetString("port")
	user, _ := cmd.Flags().GetString("user")
	pass, _ := cmd.Flags().GetString("pass")
	path, _ := cmd.Flags().GetString("path")
	databaseName, _ := cmd.Flags().GetString("database")
	amount, ok := strconv.Atoi(args[0])
	if ok != nil {
		fmt.Fprintln(os.Stderr, "The amount of documents you want needs to be the first argument")
		return
	}
	blocks, err := schema.ReadAllBlocks(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing the schemas: %v", err)
		return
	}
	database.ConnectDatabase(host, port, user, pass, databaseName)
	for _, arg := range args[1:] {
		block, err := schema.GetBlockFromName(arg, blocks)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error retreiving block: %v", err)
			return
		}
		documents := generation.GenerateDocuments(block, amount)
		database.InsertDocuments(documents, arg)
	}
}
