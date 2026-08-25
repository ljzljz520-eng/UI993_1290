package admin

import (
	"fmt"
	"refereequal/domain"
	"refereequal/repository"
)

type Service struct{ R *repository.Repository }

func New(r *repository.Repository) *Service { return &Service{R: r} }
func (s *Service) AddReferee(r domain.Referee) error {
	if e := domain.ValidateReferee(r); e != nil {
		return e
	}
	return s.R.CreateReferee(r)
}
func (s *Service) UpdateReferee(r domain.Referee) error {
	if e := domain.ValidateReferee(r); e != nil {
		return e
	}
	if _, e := s.R.GetReferee(r.ID); e != nil {
		return e
	}
	return s.R.CreateReferee(r)
}
func (s *Service) DisableReferee(id string) error {
	r, e := s.R.GetReferee(id)
	if e != nil {
		return e
	}
	r.Active = false
	return s.R.CreateReferee(r)
}
func (s *Service) AddQualification(q domain.Qualification) error {
	if e := domain.ValidateQualification(q); e != nil {
		return e
	}
	if _, e := s.R.GetReferee(q.RefereeID); e != nil {
		return fmt.Errorf("referee missing: %w", e)
	}
	return s.R.SaveQualification(q)
}
func (s *Service) DisableQualification(id string) error {
	qs, e := s.R.ListQualifications()
	if e != nil {
		return e
	}
	for _, q := range qs {
		if q.ID == id {
			q.Active = false
			return s.R.SaveQualification(q)
		}
	}
	return fmt.Errorf("qualification missing")
}
func (s *Service) AddTraining(t domain.TrainingRecord) error {
	if e := domain.ValidateTraining(t); e != nil {
		return e
	}
	return s.R.SaveTraining(t)
}
func (s *Service) RecordAudit(a domain.AuditEvent) error { return s.R.Audit(a) }
