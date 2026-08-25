package web

import (
	"net/http/httptest"
	"path/filepath"
	"refereequal/query"
	"refereequal/repository"
	"refereequal/storage"
	"testing"
)

func TestHTTPQuery(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	h := New(query.New(repository.New(s))).Handler()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/health", nil)
	h.ServeHTTP(rr, req)
	if rr.Code != 200 {
		t.Fatal(rr.Code)
	}
}
