package mdb_test

import (
	"testing"
  "encoding/json"
  "fmt"
)

func TestGetAllStocksWithSelectedHeaders(t *testing.T) {
	headers := []string{"基本分析分數", "技術分析分數", "保留盈餘增長标准分数", "基因分析標準分數"}
	data, err := db.GetAllStocksWithSelectedHeaders(headers, "test")
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	if len(data) == 0 {
		t.Error("No data returned from handler")
	} else {
		prettyData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(prettyData))
	}
}

func TestGetAllHeadersForStock(t *testing.T) {
	stockID := "1112HK-H&H國際控股"
	data, err := db.GetAllHeadersForStock(stockID, "test")
	if err != nil {
		t.Fatalf("failed to get data: %v", err)
	}

	if len(data) == 0 {
		t.Error("No data returned from handler")
	} else {
		prettyData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(prettyData))
	}
}
