package mdb

import (
	"context"
	"fmt"
  "log"
  "time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

func LoadToDB(uri string, database string, collection string, data []map[string]interface{}) error {
	// Set client options (Uniform Resource Identifier)
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("could not ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB!")

	coll := client.Database(database).Collection(collection)

	// Convert []map[string]interface{} to []interface{} for the MongoDB driver
	insertData := make([]interface{}, len(data))
	for i, v := range data {
		insertData[i] = v
	}

	// Insert data into MongoDB
	_, err = coll.InsertMany(context.TODO(), insertData)
	if err != nil {
		return fmt.Errorf("could not insert data into MongoDB: %w", err)
	}

	fmt.Println("Data inserted into MongoDB!")

	// Disconnect from MongoDB
	err = client.Disconnect(context.TODO())
	if err != nil {
		return fmt.Errorf("could not disconnect from MongoDB: %w", err)
	}

	fmt.Println("Connection to MongoDB closed.")

	return nil
}
