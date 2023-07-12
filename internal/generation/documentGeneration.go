package generation

import (
	"github.com/polBachelin/database-populator/internal/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type valueFunction func() interface{}

var fieldTypes = map[string]valueFunction{
	"string":  getRandomStringValue,
	"boolean": getRandomBoolValue,
	"object":  getRandomObjectValue,
}

func getRandomStringValue() interface{} {
	//TODO: implement
	return ""
}

func getRandomBoolValue() interface{} {
	return false
}

func getRandomObjectValue() interface{} {
	return ""
}

func GenerateDocuments(block *schema.BlockData, amount int) []interface{} {
	var documents []interface{}

	for i := 0; i < amount; i++ {
		doc := make(bson.M)
		doc["_id"] = primitive.NewObjectID()
		for _, field := range block.Fields {
			//TODO: maybe a recursive function that generates nested objects
			// if field.Type == "object" {

			// }
			value := fieldTypes[field.Type]()
			doc[field.Name] = value
		}
		documents = append(documents, doc)
	}
	return documents
}
