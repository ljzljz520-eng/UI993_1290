package repository

import (
	"refereequal/domain"
	"refereequal/storage"
)

type Repository struct{ S *storage.Store }

func New(s *storage.Store) *Repository                               { return &Repository{S: s} }
func (r *Repository) CreateReferee(v domain.Referee) error           { return r.S.SaveReferee(v) }
func (r *Repository) GetReferee(id string) (domain.Referee, error)   { return r.S.Referee(id) }
func (r *Repository) ListReferees() ([]domain.Referee, error)        { return r.S.AllReferees() }
func (r *Repository) SaveQualification(v domain.Qualification) error { return r.S.SaveQualification(v) }
func (r *Repository) ListQualifications() ([]domain.Qualification, error) {
	return r.S.AllQualifications()
}
func (r *Repository) SaveTraining(v domain.TrainingRecord) error      { return r.S.SaveTraining(v) }
func (r *Repository) ListTrainings() ([]domain.TrainingRecord, error) { return r.S.AllTrainings() }
func (r *Repository) Audit(v domain.AuditEvent) error                 { return r.S.SaveAudit(v) }
