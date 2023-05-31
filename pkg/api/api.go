package api

import (
	"net/http"
	"radarbase/pkg/mdb"
)

func StartAPIServer(db *mdb.Database) {
	h := &Handler{
		DB: db,
	}

	r := NewRouter(h)

	http.ListenAndServe(":8080", r)
}
