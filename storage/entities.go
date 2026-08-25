package storage

import (
	"os"
	"refereequal/domain"
)

func (s *Store) SaveReferee(v domain.Referee) error { return put(s, buckets, v.ID, v) }
func (s *Store) Referee(id string) (domain.Referee, error) {
	return get[domain.Referee](s, buckets, id)
}
func (s *Store) DeleteReferee(id string) error                  { return del(s, buckets, id) }
func (s *Store) AllReferees() ([]domain.Referee, error)         { return scan[domain.Referee](s, buckets) }
func (s *Store) SaveQualification(v domain.Qualification) error { return put(s, qb, v.ID, v) }
func (s *Store) Qualification(id string) (domain.Qualification, error) {
	return get[domain.Qualification](s, qb, id)
}
func (s *Store) AllQualifications() ([]domain.Qualification, error) {
	return scan[domain.Qualification](s, qb)
}
func (s *Store) SaveTraining(v domain.TrainingRecord) error { return put(s, tb, v.ID, v) }
func (s *Store) AllTrainings() ([]domain.TrainingRecord, error) {
	return scan[domain.TrainingRecord](s, tb)
}
func (s *Store) SaveAudit(v domain.AuditEvent) error     { return put(s, ab, v.ID, v) }
func (s *Store) AllAudits() ([]domain.AuditEvent, error) { return scan[domain.AuditEvent](s, ab) }
func IsNotFound(e error) bool                            { return e == os.ErrNotExist }
