package api

import (
  "fmt"
  "errors"
  "strings"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) GetByHeadersHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	collectionName := vars["collectionName"]
	headersParam, ok := r.URL.Query()["headers"]

	if !ok || len(headersParam[0]) < 1 {
		return nil, errors.New("missing headers query parameter")
	}

	headers := strings.Split(headersParam[0], ",")

	data, err := api.db.GetByHeaders(headers, collectionName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (api *API) GetSingleRecordHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	collectionName := vars["collectionName"]
	stockName := vars["stockName"]

  fmt.Println("stockName: ", stockName)

	data, err := api.db.GetSingleRecord(stockName, collectionName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (api *API) GetHeadersHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	collectionName := mux.Vars(r)["collectionName"]

	data, err := api.db.GetHeaders(collectionName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (api *API) GetMinMaxDataHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	collectionName := mux.Vars(r)["collectionName"]

	minData, maxData, err := api.db.GetMinMaxData(collectionName)
	if err != nil {
		return nil, err
	}

	return map[string]map[string]float64{
		"min": minData,
		"max": maxData,
	}, nil
}

