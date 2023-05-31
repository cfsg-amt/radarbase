package mdb

import (
	"context"
	"fmt"
  "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllStocksWithSelectedHeaders returns all the stocks with only the selected headers
func (db *Database) GetAllStocksWithSelectedHeaders(headers []string, collectionName string) ([]bson.M, error) {
	projection := bson.M{}
	for _, header := range headers {
		projection[header] = 1
	}
	opts := options.Find().SetProjection(projection)

	cursor, err := db.Collection(collectionName).Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find stocks: %w", err)
	}
	defer cursor.Close(context.Background())

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("failed to decode cursor: %w", err)
	}

	return results, nil
}

// GetAllHeadersForStock returns all the headers for a specific stock
func (db *Database) GetAllHeadersForStock(stockID string, collectionName string) (bson.M, error) {
	result := db.Collection(collectionName).FindOne(context.Background(), bson.M{"stockid": stockID})
	if result.Err() != nil {
		return nil, fmt.Errorf("failed to find stock: %w", result.Err())
	}

	var stock bson.M
	if err := result.Decode(&stock); err != nil {
		return nil, fmt.Errorf("failed to decode result: %w", err)
	}

	return stock, nil
}

// GetStockHeaders retrieves all headers from a particular collection
func (db *Database) GetStockHeaders(collectionName string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get collection for headers
	headersCollection := db.database.Collection("headers")

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
