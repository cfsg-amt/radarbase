package mdb

import (
	"context"
  "fmt"
)

func (db *Database) LoadToDB(data []map[string]interface{}) error {
	// Convert []map[string]interface{} to []interface{} for the MongoDB driver
	insertData := make([]interface{}, len(data))
	for i, v := range data {
		insertData[i] = v
	}

	// Insert data into MongoDB
	_, err := db.collection.InsertMany(context.Background(), insertData)
	if err != nil {
		return fmt.Errorf("could not insert data into MongoDB: %w", err)
	}

	fmt.Println("Data inserted into MongoDB!")

	return nil
}
