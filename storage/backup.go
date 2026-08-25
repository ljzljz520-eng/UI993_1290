package storage

import (
	"encoding/json"
	"refereequal/domain"
)

func (s *Store) Export() ([]byte, error) {
	r, e := s.AllReferees()
	if e != nil {
		return nil, e
	}
	q, e := s.AllQualifications()
	if e != nil {
		return nil, e
	}
	t, e := s.AllTrainings()
	if e != nil {
		return nil, e
	}
	return json.Marshal(struct {
		Referees       []domain.Referee
		Qualifications []domain.Qualification
		Trainings      []domain.TrainingRecord
	}{r, q, t})
}
func (s *Store) Import(data []byte) error {
	var x struct {
		Referees       []domain.Referee
		Qualifications []domain.Qualification
		Trainings      []domain.TrainingRecord
	}
	if e := json.Unmarshal(data, &x); e != nil {
		return e
	}
	for _, r := range x.Referees {
		if e := s.SaveReferee(r); e != nil {
			return e
		}
	}
	for _, q := range x.Qualifications {
		if e := s.SaveQualification(q); e != nil {
			return e
		}
	}
	for _, t := range x.Trainings {
		if e := s.SaveTraining(t); e != nil {
			return e
		}
	}
	return nil
}
