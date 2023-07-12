package database

import (
	"context"
)

func InsertDocuments(documents []interface{}, collectionName string) error {
	collection := databaseConnection.db.Collection(collectionName)

	_, err := collection.InsertMany(context.TODO(), documents)
	return err
}
