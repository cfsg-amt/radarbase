package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"radarbase/pkg/mdb"
)

type Server struct {
	db *mdb.Database
}

func NewServer(db *mdb.Database) *Server {
	return &Server{db: db}
}

func (s *Server) HandleGetAllStocksWithSelectedHeaders(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	headers := strings.Split(params.Get("headers"), ",")

	data, err := s.db.GetAllStocksWithSelectedHeaders(headers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (s *Server) HandleGetAllHeadersForStock(w http.ResponseWriter, r *http.Request) {
	stockID := strings.TrimPrefix(r.URL.Path, "/stocks/")

	data, err := s.db.GetAllHeadersForStock(stockID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(data)
}
