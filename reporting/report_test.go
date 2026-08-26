package reporting

import (
	"refereequal/domain"
	"testing"
)

func TestSummary(t *testing.T) {
	v := Summary(domain.QueryResult{Referee: domain.Referee{Name: "N"}})
	if v == "" {
		t.Fatal()
	}
	if ActiveCount([]domain.Qualification{{Active: true}}) != 1 {
		t.Fatal()
	}
}
