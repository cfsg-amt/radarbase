package api

import (
	"encoding/json"
	"net/http"

	"radarbase/pkg/mdb"
)

type API struct {
	db *mdb.MDB
}

func NewAPI(db *mdb.MDB) *API {
	return &API{db: db}
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
