package api_test

import (
	"context"
	"fmt"
  "time"
	"log"
	"net/http"
	"os"
	"testing"

	"radarbase/pkg/api"
	"radarbase/pkg/mdb"
	"radarbase/pkg/excel"
)

var db *mdb.MDB
var ctx context.Context

var srv *http.Server

func TestMain(m *testing.M) {
	var err error
	db, err = mdb.NewMDB("mongodb://localhost:27017", "testRowDB", "testColDB")
	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	// Parse row data
	rowData, rowHeaders, err := excel.RowParse("testdata/sample.xlsx", "Sheet1")
  fmt.Println(rowHeaders) // TODO: remove this

	if err != nil {
    fmt.Println(err)
		os.Exit(1)
	}

	// Parse columnar data
	colData, colHeaders, err := excel.ColParse("testdata/sample.xlsx", "Sheet1")
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

	// Initialize API
	api := api.NewAPI(db)

	// Set up HTTP server
	srv = &http.Server{
		Handler: api.SetupRouter(),
		Addr:    "127.0.0.1:8081",
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			// it's fine to panic here, as this should never happen when closing the server
			log.Panic(err)
		}
	}()

	// Running all the tests
	code := m.Run()

	// Cleanup
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

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
