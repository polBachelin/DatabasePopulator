package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates documents for a MongoDB database",
	Long:  ``,
	Run:   generate,
}

type database struct {
	db *mongo.Database
}

var databaseConnection = &database{
	db: nil,
}

type FileData struct {
	Blocks []BlockData `yaml:"blocks"`
}

type BlockData struct {
	Name   string      `yaml:"name"`
	Fields []FieldData `yaml:"fields"`
}

type FieldData struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func connectDatabase(host, port, user, pass, dbName string) *mongo.Database {
	uri := "mongodb://" + user + ":" + pass + "@" + host + ":" + port
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	databaseConnection.db = client.Database(dbName)
	return databaseConnection.db
}

func ReadBlockFile(filename string) (*FileData, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &FileData{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("error in file %s: %v", filename, err)
	}
	return c, err
}

func ReadAllBlocks(directory string) ([]*FileData, error) {
	entries, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error in directory: %v", err)
		return nil, err
	}
	data := make([]*FileData, 0, len(entries))
	for _, e := range entries {
		block, err := ReadBlockFile(filepath.Join(directory, e.Name()))
		if err != nil {
			return data, err
		}
		data = append(data, block)
	}
	return data, err
}

func generate(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetString("host")
	port, _ := cmd.Flags().GetString("port")
	user, _ := cmd.Flags().GetString("user")
	pass, _ := cmd.Flags().GetString("pass")
	//path, _ := cmd.Flags().GetString("path")
	database, _ := cmd.Flags().GetString("database")
	db := connectDatabase(host, port, user, pass, database)

	//TODO: Generate documents here
	if len(args) < 0 {
		fmt.Fprintln(os.Stderr, "No arguments provided.")
		return
	}

}
