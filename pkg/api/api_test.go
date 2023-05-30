package api_test

import (
	"context"
	"os"
	"testing"
	"time"

	"radarbase/pkg/mdb"
	"radarbase/pkg/excel"
  "radarbase/pkg/api"
)

var server *api.Server

func TestMain(m *testing.M) {
	// Create a new database and load data into it
	db, err := mdb.NewDatabase("mongodb://localhost:27017", "testdb", "stocks")
	if err != nil {
		os.Exit(1)
	}

	data, err := excel.Parse("testdata/sample.xlsx", "Sheet1")
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, item := range data {
		if err = db.InsertOne(ctx, item); err != nil {
			os.Exit(1)
		}
	}

	server = api.NewServer(db)

	// Running all the tests
	code := m.Run()

	// Dropping the test database
	if err := db.Drop(ctx); err != nil {
		os.Exit(1)
	}

	// Disconnecting from MongoDB
	if err := db.Disconnect(ctx); err != nil {
		os.Exit(1)
	}

	os.Exit(code)
}
