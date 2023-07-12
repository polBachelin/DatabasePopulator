package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type database struct {
	db *mongo.Database
}

var databaseConnection = &database{
	db: nil,
}

func ConnectDatabase(host, port, user, pass, dbName string) *mongo.Database {
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
