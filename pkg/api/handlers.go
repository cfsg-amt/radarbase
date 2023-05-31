package api

import (
  "encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"radarbase/pkg/mdb"
  "strings"
)

// handler struct contains the dependencies for the handlers
type Handler struct {
	DB *mdb.Database
}

// GetStocksHandler returns all stocks from a particular collection
func (h *Handler) GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	// extract the collection name from the URL path variable
	vars := mux.Vars(r)
	collectionName := vars["collectionName"]

	// Get headers from query params
	headersParam := r.URL.Query()["headers"]
	if len(headersParam) == 0 {
		http.Error(w, "Missing headers query param", http.StatusBadRequest)
		return
	}
	headers := strings.Split(headersParam[0], ",")

	// fetch stocks with selected headers from the database
	data, err := h.DB.GetAllStocksWithSelectedHeaders(headers, collectionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write data as JSON to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// GetStockByIDHandler returns a particular stock from a collection
func (h *Handler) GetStockByIDHandler(w http.ResponseWriter, r *http.Request) {
	// extract the collection name and stock ID from the URL path variables
	vars := mux.Vars(r)
	collectionName := vars["collectionName"]
	stockID := vars["stockID"]

	// fetch the stock with the given ID from the database
	stock, err := h.DB.GetAllHeadersForStock(stockID, collectionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Write data as JSON to response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stock)
}
