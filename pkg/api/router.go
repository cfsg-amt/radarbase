package api

import (
	"github.com/gorilla/mux"
)

func NewRouter(h *Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/{collectionName}/stocks", h.GetStocksHandler).Methods("GET")
	r.HandleFunc("/api/v1/{collectionName}/stocks/{stockID}", h.GetStockByIDHandler).Methods("GET")

	return r
}
