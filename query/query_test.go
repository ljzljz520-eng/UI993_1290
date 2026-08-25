package query

import (
	"path/filepath"
	"refereequal/domain"
	"refereequal/repository"
	"refereequal/storage"
	"testing"
)

func TestWorkflowQuery(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := repository.New(s)
	r.CreateReferee(domain.Referee{ID: "r", Name: "N", Active: true})
	r.SaveQualification(domain.Qualification{ID: "q", RefereeID: "r", Level: "A", Sport: "Run", ValidUntil: "2027-01-01", Active: true})
	v, e := New(r).Find("r", "n")
	if e != nil || len(v.Qualifications) != 1 {
		t.Fatal(v, e)
	}
}
