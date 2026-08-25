package admin

import (
	"path/filepath"
	"refereequal/domain"
	"refereequal/repository"
	"refereequal/storage"
	"testing"
)

func TestWorkflowAdmin(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := repository.New(s)
	a := New(r)
	if e := a.AddReferee(domain.Referee{ID: "r", Name: "N", Active: true}); e != nil {
		t.Fatal(e)
	}
	if e := a.AddQualification(domain.Qualification{ID: "q", RefereeID: "r", Level: "A", Sport: "Run", ValidUntil: "2027-01-01", Active: true}); e != nil {
		t.Fatal(e)
	}
	if e := a.DisableQualification("q"); e != nil {
		t.Fatal(e)
	}
}
