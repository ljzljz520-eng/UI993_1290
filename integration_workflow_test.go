package referequal

import (
	"path/filepath"
	"refereequal/admin"
	"refereequal/domain"
	"refereequal/query"
	"refereequal/repository"
	"refereequal/storage"
	"testing"
)

func TestBusinessChain13(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r := repository.New(s)
	a := admin.New(r)
	a.AddReferee(domain.Referee{ID: "r", Name: "N", Active: true})
	a.AddQualification(domain.Qualification{ID: "b", RefereeID: "r", Level: "A", Sport: "Run", ValidUntil: "2027-01-01", Active: true})
	a.AddQualification(domain.Qualification{ID: "a", RefereeID: "r", Level: "A", Sport: "Run", ValidUntil: "2027-01-01", Active: true})
	v, e := query.New(r).Find("r", "N")
	if e != nil {
		t.Fatal(e)
	}
	if len(v.Qualifications) != 2 || v.Qualifications[0].ID != "b" {
		t.Fatalf("unstable tie order: %#v", v.Qualifications)
	}
}
