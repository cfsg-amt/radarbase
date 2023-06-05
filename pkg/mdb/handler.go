package mdb

import (
	"context"
	"fmt"
  "time"
  "encoding/hex"
  "crypto/sha256"
  
	"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *MDB) GetSingleRecord(name string, collectionName string) (bson.M, error) {
  // Generate SHA256 hash of the stockName
  hash := sha256.New()
  hash.Write([]byte(name))
  hashStr := hex.EncodeToString(hash.Sum(nil))
  result := db.RowCollection(collectionName).FindOne(context.Background(), bson.M{"_id": hashStr})

  if result.Err() != nil {
    return nil, fmt.Errorf("failed to find stock: %w", result.Err())
  }

	var stock bson.M
	if err := result.Decode(&stock); err != nil {
		return nil, fmt.Errorf("failed to decode result: %w", err)
	}

	return stock, nil
}

func (db *MDB) GetHeaders(collectionName string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get collection for headers
	headersCollection := db.coldb.Collection("headers")

	// Find the headers document
	var result struct {
		Headers []string `bson:"headers"`
	}

	err := headersCollection.FindOne(ctx, bson.M{"_id": collectionName}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Headers, nil
}


func (db *MDB) GetByHeaders(headers []string, collectionName string) (map[string]map[string][]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// The map to hold grouped results
	groupedResult := make(map[string]map[string][]interface{})

	// The slice to hold not found headers
	notFoundHeaders := []string{}

	collection := db.ColCollection(collectionName)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Load all documents into memory
	var documents []bson.M
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	for _, header := range headers {
		headerFound := false

		for _, document := range documents {
			// Check if header exists in the document
			if values, ok := document[header]; ok {
				headerFound = true

				// If the group does not exist in the result map, create it
				if _, ok := groupedResult[header]; !ok {
					groupedResult[header] = make(map[string][]interface{})
				}

				// Append the values to the result map
        // primitive.A in Go is an alias for []interface{}
        groupedResult[header][document["_id"].(string)] = values.(primitive.A)
			}
		}

		// If the header was not found in any documents, add it to notFoundHeaders
		if !headerFound {
			notFoundHeaders = append(notFoundHeaders, header)
		}
	}

	// Check for errors from iterating over documents.
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(notFoundHeaders) > 0 {
		return groupedResult, fmt.Errorf("the following headers %v were not found in any collections", notFoundHeaders)
	}

	return groupedResult, nil
}
