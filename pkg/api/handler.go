package api

import (
  "fmt"
  "errors"
  "strings"
	"net/http"
  "go.etcd.io/bbolt"
  "golang.org/x/crypto/bcrypt"
  "encoding/json"

	"github.com/gorilla/mux"
)

type PostRequest struct {
    AdminPwd string `json:"adminpwd"`
    Key      string `json:"key"`
    Value    string `json:"value"`
}

func (api *API) SetKeyValueHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    var request PostRequest
    err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        return nil, err
    }

    // Hardcoded hashed password
    hashedAdminPwd := []byte("$2y$10$v4Iqpj1aHirivVOseCnMiuyR9Tf1qC7ciom.wg2As0Z1DG.hvtcce")  

    // Compare provided password with hashed password
    err = bcrypt.CompareHashAndPassword(hashedAdminPwd, []byte(request.AdminPwd))
    if err != nil {
        return nil, fmt.Errorf("invalid admin password")  // Don't give too much detail in the error message
    }

    err = api.SetValue([]byte(request.Key), []byte(request.Value))
    if err != nil {
        return nil, err
    }

    return "Value set successfully", nil
}

func (api *API) SetValue(key, value []byte) error {
    return api.kv.Update(func(tx *bbolt.Tx) error {
	    b, err := tx.CreateBucketIfNotExists([]byte("login"))
	    if err != nil {
		    return err
	    }

	    return b.Put(key, value)
    })
}

func (api *API) GetValue(key []byte) ([]byte, error) {
    var value []byte
    err := api.kv.View(func(tx *bbolt.Tx) error {
	    b := tx.Bucket([]byte("login"))
	    if b == nil {
		    return fmt.Errorf("bucket %q not found", "login")
	    }

	    value = b.Get(key)
	    return nil
    })

    return value, err
}

func (api *API) GetKeyValueHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
    vars := mux.Vars(r)
    key := vars["key"]

    value, err := api.GetValue([]byte(key))
    if err != nil {
        return nil, err
    }

    return string(value), nil
}

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

