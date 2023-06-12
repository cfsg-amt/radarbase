package api

import (
  "encoding/json"
  "net/http"

	"github.com/gorilla/mux"
)

func (api *API) SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/{collectionName}/item", api.handleRequest(api.GetByHeadersHandler)).Methods("GET")
	router.HandleFunc("/api/v1/{collectionName}/item/{stockName}", api.handleRequest(api.GetSingleRecordHandler)).Methods("GET")
	router.HandleFunc("/api/v1/headers/{collectionName}", api.handleRequest(api.GetHeadersHandler)).Methods("GET")
  router.HandleFunc("/api/v1/minmax/{collectionName}", api.handleRequest(api.GetMinMaxDataHandler)).Methods("GET")

	return router
}

// Middleware function to handle common logic and error handling
func (api *API) handleRequest(handler func(w http.ResponseWriter, r *http.Request) (interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := handler(w, r)
		if err != nil {
			api.respondWithError(w, http.StatusInternalServerError, err)
			return
		}

    // Set CORS headers (just for development) TODO: remove this
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
