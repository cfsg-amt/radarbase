package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"radarbase/pkg/mdb"
)

type Handler struct {
	DB *mdb.Database
}

func (h *Handler) GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	// To be implemented
}

func (h *Handler) GetStockByIDHandler(w http.ResponseWriter, r *http.Request) {
	// To be implemented
}
