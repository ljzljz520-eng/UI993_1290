package repository

import (
	"path/filepath"
	"refereequal/domain"
	"refereequal/storage"
	"testing"
)

func TestRepositoryRoundTrip(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := New(s)
	r.CreateReferee(domain.Referee{ID: "a", Name: "A"})
	v, e := r.GetReferee("a")
	if e != nil || v.ID != "a" {
		t.Fatal(v, e)
	}
}
