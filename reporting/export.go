package reporting

import (
	"encoding/json"
	"refereequal/domain"
)

func JSON(r domain.QueryResult) ([]byte, error) { return json.Marshal(r) }
func Levels(qs []domain.Qualification) map[string]int {
	m := map[string]int{}
	for _, q := range qs {
		m[q.Level]++
	}
	return m
}
