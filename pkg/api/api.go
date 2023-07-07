package api

import (
	"encoding/json"
	"net/http"
  "log"

  "go.etcd.io/bbolt"
	"radarbase/pkg/mdb"
)

type API struct {
	db *mdb.MDB
  kv *bbolt.DB
}

func NewAPI(db *mdb.MDB) *API {
  kv, err := bbolt.Open("kv.db", 0600, nil)
  // 0600 means "only allow the owner of this file to read and write it".
  if err != nil {
    log.Fatal(err)
  }
	return &API{db: db, kv: kv}
}

// writeJSON is a helper function for writing a response as JSON.
func (api *API) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// respondWithError is a helper function to respond with an error.
func (api *API) respondWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	api.writeJSON(w, map[string]string{"error": err.Error()})
}
