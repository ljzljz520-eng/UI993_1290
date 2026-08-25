package domain

import "testing"

func TestDates(t *testing.T) {
	if !ValidDate("2026-01-01") || ValidDate("bad") {
		t.Fatal()
	}
	q := Qualification{Active: true, ValidUntil: "2027-01-01"}
	if !q.IsValid("2026-01-01") {
		t.Fatal()
	}
}
