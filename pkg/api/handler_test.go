package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestGetByHeadersHandler(t *testing.T) {
	headers := []string{"基本分析分數", "技術分析分數", "保留盈餘增長标准分数", "基因分析標準分數"}
	encodedHeaders := url.QueryEscape(strings.Join(headers, ","))

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/v1/test/item?headers=%s", encodedHeaders))

	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status: got %v want %v", resp.Status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	// Print the raw JSON response
	t.Logf("JSON Response: %s\n", body)

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	// Print the unmarshaled response
	t.Logf("Map Response: %+v\n", data)
}


func TestGetSingleRecordHandler(t *testing.T) {
	stockName := "1112HK-H&H國際控股"
	encodedStockName := url.QueryEscape(stockName)

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/v1/test/item/%s", encodedStockName))

	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status: got %v want %v", resp.Status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	// Print the raw JSON response
	t.Logf("JSON Response: %s\n", body)

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	// Print the unmarshaled response
	t.Logf("Map Response: %+v\n", data)
}


func TestGetHeadersHandler(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/v1/headers/test")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status: got %v want %v", resp.Status, http.StatusOK)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	// Print the raw JSON response
	t.Logf("JSON Response: %s\n", body)

	var data []string
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	// Print the unmarshaled response
	t.Logf("Headers: %+v\n", data)
}
