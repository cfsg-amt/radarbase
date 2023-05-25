package mdb

import (
	"testing"
  "os" 
	"radarbase/pkg/excel"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func TestMain(m *testing.M) {
    var err error
    client, err = ConnectToDB() 
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Creating test database
    testDB := client.Database("testdb")

    // Call all tests
    code := m.Run()

    // Delete test database
    err = testDB.Drop(context.Background())
    if err != nil {
        log.Fatal("Failed to drop test database: ", err)
    }

    // Disconnect from MongoDB
    if err := client.Disconnect(context.Background()); err != nil {
        log.Fatalf("Failed to disconnect from MongoDB: %v", err)
    }

    os.Exit(code)
}

func TestInsertToDB(t *testing.T) {
	// Parse excel file
	data, err := excel.Parse("testdata/sample.xlsx", "Sheet1")
	if err != nil {
		t.Fatalf("Parse failed with error: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		t.Fatalf("Could not ping MongoDB: %v", err)
	}

	// Connect to the test database and collection
	collection := client.Database("testdb").Collection("stocks")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	for _, item := range data {
		_, err := collection.InsertOne(ctx, item)
		if err != nil {
			t.Errorf("Insert failed with error: %v", err)
		}
	}

}
