package reporting

import (
	"fmt"
	"refereequal/domain"
)

func Summary(r domain.QueryResult) string {
	return fmt.Sprintf("%s:%d qualifications/%d trainings", r.Referee.Name, len(r.Qualifications), len(r.Trainings))
}
func QualificationLabels(qs []domain.Qualification) []string {
	out := make([]string, 0, len(qs))
	for _, q := range qs {
		out = append(out, q.Level+" "+q.Sport)
	}
	return out
}
func ActiveCount(qs []domain.Qualification) int {
	n := 0
	for _, q := range qs {
		if q.Active {
			n++
		}
	}
	return n
}
