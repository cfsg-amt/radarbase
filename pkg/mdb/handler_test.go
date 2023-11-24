package mdb_test

import (
	"testing"
  "encoding/json"
  "fmt"
)

func TestGetMinMaxData(t *testing.T) {
  print("TestGetMinMaxData\n")
	minData, maxData, err := db.GetMinMaxData("test")
	if err != nil {
		t.Fatalf("failed to get min/max data: %v", err)
	}

	fmt.Println("Min data:")
	for header, minValue := range minData {
		fmt.Printf("Header: %s, Min Value: %f\n", header, minValue)
	}

	fmt.Println("Max data:")
	for header, maxValue := range maxData {
		fmt.Printf("Header: %s, Max Value: %f\n", header, maxValue)
	}
}

func TestGetHeaders(t *testing.T) {
  print("TestGetHeaders\n")
	// Headers to request
	headers := []string{"基本分析分數", "技術分析分數", "基因分析標準分數", "name"}

	// Call the GetHeaders function
	result, err := db.GetByHeaders(headers, "test")
	if err != nil {
		t.Fatalf("failed to get headers: %v", err)
	}

	// Check if the result is empty
	if len(result) == 0 {
		t.Error("No headers returned from handler")
	} else {
		// Print out the result in a pretty format
		prettyResult, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(prettyResult))
	}
}

func TestGetAllHeadersForStock(t *testing.T) {
  print("TestGetAllHeadersForStock\n")
  // stockName := "1070HK-ＴＣＬ電子"
  stockName := "消費奢侈品"

	data, err := db.GetSingleRecord(stockName, "test")


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

func TestGetStockHeaders(t *testing.T) {
	headers, err := db.GetHeaders("test")
	if err != nil {
		t.Fatalf("failed to get headers: %v", err)
	}

	if len(headers) == 0 {
		t.Error("No headers returned from handler")
	} else {
		fmt.Println(headers)
	}
}
