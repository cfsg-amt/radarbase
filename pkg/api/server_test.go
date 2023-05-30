package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
  "context"
  "fmt"

	"radarbase/pkg/mdb"
	"radarbase/pkg/api"
)

func TestHandleGetAllStocksWithSelectedHeaders(t *testing.T) {
	// create an instance of your database and load some test data into it
	db, _ := mdb.NewDatabase("mongodb://localhost:27017", "testdb", "stocks")
	defer db.Disconnect(context.Background())

	// create an instance of your server with the test database
	server := api.NewServer(db)

	// create a new HTTP request
	req, err := http.NewRequest("GET", "/stocks?headers=基本分析分數, 技術分析分數, 基因分析標準分數", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// call the handler function
	handler := http.HandlerFunc(server.HandleGetAllStocksWithSelectedHeaders)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

  fmt.Println(rr.Body)
}

func TestHandleGetAllHeadersForStock(t *testing.T) {
	// create an instance of your database and load some test data into it
	db, _ := mdb.NewDatabase("mongodb://localhost:27017", "testdb", "stocks")
	defer db.Disconnect(context.Background())

	// create an instance of your server with the test database
	server := api.NewServer(db)

	// create a new HTTP request
	req, err := http.NewRequest("GET", "/stocks/1112HK-H&H國際控股", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// call the handler function
	handler := http.HandlerFunc(server.HandleGetAllHeadersForStock)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

  fmt.Println(rr.Body)
}
