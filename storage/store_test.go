package storage

import (
	"path/filepath"
	"refereequal/domain"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveReferee(domain.Referee{ID: "x", Name: "X", Active: true}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, e := s.Referee("x")
	if e != nil || r.Name != "X" {
		t.Fatal(r, e)
	}
}
