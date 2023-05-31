package mdb_test

import (
	"context"
	"os"
	"testing"
	"time"
  "fmt"

	"radarbase/pkg/mdb"
	"radarbase/pkg/excel"
)

var db *mdb.Database

func TestMain(m *testing.M) {
	var err error
	db, err = mdb.NewDatabase("mongodb://localhost:27017", "testdb")
	if err != nil {
		os.Exit(1)
	}

	data, headers, err := excel.Parse("testdata/sample.xlsx", "Sheet1")
  fmt.Println(headers)

	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

  // Inserting the test data
  db.LoadToDB(data, headers, "test")

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
