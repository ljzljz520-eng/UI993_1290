package query

import (
	"refereequal/domain"
	"refereequal/repository"
	"sort"
	"strings"
)

type Service struct{ R *repository.Repository }

func New(r *repository.Repository) *Service { return &Service{R: r} }
func (s *Service) Find(id, name string) (domain.QueryResult, error) {
	r, e := s.R.GetReferee(id)
	if e != nil {
		return domain.QueryResult{}, e
	}
	if name != "" && !strings.EqualFold(r.Name, name) {
		return domain.QueryResult{}, repositoryErr("name mismatch")
	}
	qs, _ := s.R.ListQualifications()
	ts, _ := s.R.ListTrainings()
	out := domain.QueryResult{Referee: r}
	for _, q := range qs {
		if q.RefereeID == id && q.Active {
			out.Qualifications = append(out.Qualifications, q)
		}
	}
	for _, t := range ts {
		if t.RefereeID == id {
			out.Trainings = append(out.Trainings, t)
		}
	}
	sort.Slice(out.Qualifications, func(i, j int) bool { return out.Qualifications[i].Level < out.Qualifications[j].Level })
	sort.Slice(out.Trainings, func(i, j int) bool { return out.Trainings[i].Score > out.Trainings[j].Score })
	return out, nil
}

type repositoryErr string

func (e repositoryErr) Error() string { return string(e) }
