package api_test

import (
	"context"
  "strings"
  "net/url"
	"fmt"
  "time"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"radarbase/pkg/api"
	"radarbase/pkg/mdb"
	"radarbase/pkg/excel"
)

var db *mdb.Database
var ts *httptest.Server

func TestMain(m *testing.M) {
	// Setup function
	var err error
	db, err = mdb.NewDatabase("mongodb://localhost:27017", "testdb")
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
		if err = db.InsertOne(ctx, "test", item); err != nil {
			os.Exit(1)
		}
	}

	defer db.Disconnect(context.Background())

	handler := &api.Handler{DB: db}

  r := api.NewRouter(handler)

	ts = httptest.NewServer(r)
	defer ts.Close()

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

	// Exit
	os.Exit(code)
}

func TestAPI_1(t *testing.T) {
  // headers
  headers := []string{"基本分析分數", "技術分析分數", "保留盈餘增長标准分数", "基因分析標準分數"}

  // Construct headers query param
  headersParam := url.QueryEscape(strings.Join(headers, ","))

  // Test GetStocksHandler
  res, err := http.Get(fmt.Sprintf("%s/api/v1/test/stocks?headers=%s", ts.URL, headersParam))
  if err != nil {
      log.Fatal(err)
  }

  stocks, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("%s\n", stocks)
}

func TestAPI_2(t *testing.T) {
  // stockID
  stockID := "1112HK-H&H國際控股"
  // Test GetStockByIDHandler
  res, err := http.Get(fmt.Sprintf("%s/api/v1/test/stocks/%s", ts.URL, url.QueryEscape(stockID)))
  if err != nil {
      log.Fatal(err)
  }

  stock, err := ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err != nil {
      log.Fatal(err)
  }

  fmt.Printf("%s", stock)
}
