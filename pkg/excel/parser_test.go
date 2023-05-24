package excel

import (
	"testing"
	"reflect"
)

func TestParse(t *testing.T) {
	data, err := Parse("testdata/sample.xlsx", "Sheet1")
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
	
	// Print the parsed data for visual inspection
	// This is not generally part of testing but can help verify that the parsing is done correctly
	for _, item := range data {
    t.Logf("Parsed data: %+v", item)
	}
}
