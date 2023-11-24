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

var db *mdb.MDB

func TestMain(m *testing.M) {
	var err error
	db, err = mdb.NewMDB("mongodb://localhost:27017", "testRowDB", "testColDB")
	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	// Parse row data
	rowData, rowHeaders, err := excel.RowParse("testdata/sample1.xlsx", "Sheet1")
  fmt.Println(rowHeaders) // TODO: remove this

	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	// Parse columnar data
	colData, colHeaders, err := excel.ColParse("testdata/sample1.xlsx", "Sheet1")
  fmt.Println(colHeaders)

	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

  // Create test collection
  // Inserting the test row data
  err = db.RowLoadToDB(rowData, "test")
  if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	// Inserting the test columnar data
  err = db.ColLoadToDB(colData, colHeaders, "test")
  if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	// Running all the tests
	code := m.Run()

	// Dropping the test databases
	if err := db.Drop(ctx); err != nil {
    fmt.Println(err)
    os.Exit(1)
	}

	// Disconnecting from MongoDB
	if err := db.Disconnect(ctx); err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(code)
}
