package web

import (
	"encoding/json"
	"net/http"
	"refereequal/query"
)

type API struct{ Q *query.Service }

func New(q *query.Service) *API { return &API{Q: q} }
func (a *API) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK); w.Write([]byte("ok")) })
	mux.HandleFunc("/query", a.query)
	return mux
}
func (a *API) query(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	name := r.URL.Query().Get("name")
	if id == "" {
		http.Error(w, "id required", 400)
		return
	}
	v, e := a.Q.Find(id, name)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
