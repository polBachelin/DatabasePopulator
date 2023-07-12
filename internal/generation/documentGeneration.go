package generation

import (
	"github.com/polBachelin/database-populator/internal/schema"
	"go.mongodb.org/mongo-driver/bson"
)

type valueFunction func() interface{}

var fieldTypes = map[string]valueFunction{
	"string": getRandomStringValue,
}

func getRandomStringValue() interface{} {
	//TODO: implement
	return ""
}

func GenerateDocuments(block *schema.BlockData, amount int) []interface{} {
	var documents []interface{}

	for i := 0; i < amount; i++ {
		for _, field := range block.Fields {
			//TODO: maybe a recursive function that generates nested objects
			// if field.Type == "object" {

			// }
			value := fieldTypes[field.Type]()
			doc := bson.M{field.Name: value}
			documents = append(documents, doc)
		}
	}
	return documents
}
