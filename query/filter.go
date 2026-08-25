package query

import (
	"refereequal/domain"
	"sort"
)

func FilterBySport(in []domain.Qualification, sport string) []domain.Qualification {
	out := []domain.Qualification{}
	for _, q := range in {
		if q.Sport == sport {
			out = append(out, q)
		}
	}
	return out
}
func SortByValidity(in []domain.Qualification) []domain.Qualification {
	out := append([]domain.Qualification{}, in...)
	sort.Slice(out, func(i, j int) bool {
		if out[i].ValidUntil == out[j].ValidUntil {
			return out[i].ID < out[j].ID
		}
		return out[i].ValidUntil < out[j].ValidUntil
	})
	return out
}
func RecentTrainings(in []domain.TrainingRecord, limit int) []domain.TrainingRecord {
	out := append([]domain.TrainingRecord{}, in...)
	sort.Slice(out, func(i, j int) bool { return out[i].CompletedOn > out[j].CompletedOn })
	if limit > 0 && len(out) > limit {
		return out[:limit]
	}
	return out
}
