package excel_test

import (
	"testing"
	"reflect"
  "fmt"
  "radarbase/pkg/excel"
)

func TestRowParse(t *testing.T) {
	data, headers, err := excel.RowParse("testdata/sample.xlsx", "Sheet1")
	if err != nil {
		t.Errorf("Parse failed with error: %v", err)
	}
	if len(data) == 0 {
		t.Errorf("Parse returned empty data")
	}
	
	// check if the returned data is of type []map[string]interface{}
	if reflect.TypeOf(data).String() != "[]map[string]interface {}" {
		t.Errorf("Parse did not return the expected type of []map[string]interface{}")
	}

  fmt.Printf("headers: %v\n", headers)
	
	// Print the parsed data for visual inspection
	// This is not generally part of testing but can help verify that the parsing is done correctly
	for i, item := range data {
    t.Logf("Parsed data: %+v", item)
    if record, ok := item["股票"]; ok {
      t.Logf("股票 of %dth record: %v", i, record)
    }
	}
}

func TestColParse(t *testing.T) {
	data, headers, err := excel.ColParse("testdata/sample1.xlsx", "Sheet1")
	if err != nil {
		t.Errorf("ColParse failed with error: %v", err)
	}
	if len(data) == 0 {
		t.Errorf("ColParse returned empty data")
	}
	
	// check if the returned data is of type map[string][]interface{}
	if reflect.TypeOf(data).String() != "map[string][]interface {}" {
		t.Errorf("ColParse did not return the expected type of map[string][]interface{}")
	}

  fmt.Printf("headers: %v\n", headers)
	
	// Print the parsed data for visual inspection
	// This is not generally part of testing but can help verify that the parsing is done correctly
	for header, values := range data {
    t.Logf("Parsed data for header %s: %+v", header, values)
    fmt.Printf("Parsed data for header %s: %+v\n", header, values)
  }
}
