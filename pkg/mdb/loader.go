package mdb

import (
	"context"
  "fmt"

  "go.mongodb.org/mongo-driver/bson"
)

func (db *Database) LoadToDB(data []map[string]interface{}, headers []string, collectionName string) error {
	// Convert []map[string]interface{} to []interface{} for the MongoDB driver
	insertData := make([]interface{}, len(data))
	for i, v := range data {
		insertData[i] = v
	}

	// Get collection for headers
	headersCollection := db.database.Collection("headers")

	// Insert headers into MongoDB headers collection with collection name as identifier
	headerDoc := bson.M{"_id": collectionName, "headers": headers}
	_, err := headersCollection.InsertOne(context.Background(), headerDoc)
	if err != nil {
		return fmt.Errorf("could not insert headers into MongoDB: %w", err)
	}

	// Get collection for data
	dataCollection := db.database.Collection(collectionName)

	// Insert data into MongoDB
	_, err = dataCollection.InsertMany(context.Background(), insertData)
	if err != nil {
		return fmt.Errorf("could not insert data into MongoDB: %w", err)
	}

	fmt.Println("Data inserted into MongoDB!")

	return nil
}
