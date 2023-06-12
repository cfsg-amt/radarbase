package mdb

import (
	"context"
	"fmt"
  "time"
  "strconv"
  "encoding/hex"
  "crypto/sha256"
  
	"go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *MDB) GetMinMaxData(collectionName string) (minData map[string]float64, maxData map[string]float64, err error) {
	// Get collection for data
	dataCollection := db.ColCollection(collectionName)

	// Get min and max data from the database
	minMaxDoc := bson.M{}
	err = dataCollection.FindOne(context.Background(), bson.M{"_id": collectionName + "_min_max"}).Decode(&minMaxDoc)
	if err != nil {
		return nil, nil, fmt.Errorf("could not retrieve min/max data from MongoDB: %w", err)
	}

  // Convert interface{} to map[string]interface{}
  minDataPrimitiveM, ok := minMaxDoc["min"].(primitive.M)
  if !ok {
    return nil, nil, fmt.Errorf("unable to convert min data to primitive.M")
  }

  minDataInterface := make(map[string]interface{})
  bsonBytes, _ := bson.Marshal(minDataPrimitiveM)
  err = bson.Unmarshal(bsonBytes, &minDataInterface)
  if err != nil {
    return nil, nil, fmt.Errorf("unable to unmarshal min data BSON: %w", err)
  }

  minData = make(map[string]float64)
  for key, value := range minDataInterface {
    switch v := value.(type) {
    case float64:
      minData[key] = v
    case primitive.Decimal128:
      f, _ := strconv.ParseFloat(v.String(), 64)
      minData[key] = f
  }
  }

  maxDataPrimitiveM, ok := minMaxDoc["max"].(primitive.M)
  if !ok {
    return nil, nil, fmt.Errorf("unable to convert max data to primitive.M")
  }

  maxDataInterface := make(map[string]interface{})
  bsonBytes, _ = bson.Marshal(maxDataPrimitiveM)
  err = bson.Unmarshal(bsonBytes, &maxDataInterface)
  if err != nil {
    return nil, nil, fmt.Errorf("unable to unmarshal max data BSON: %w", err)
  }

  maxData = make(map[string]float64)
  for key, value := range maxDataInterface {
    switch v := value.(type) {
    case float64:
      maxData[key] = v
    case primitive.Decimal128:
      f, _ := strconv.ParseFloat(v.String(), 64)
      maxData[key] = f
  }
  }

  return minData, maxData, nil
}

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
