package admin

import (
	"fmt"
	"refereequal/domain"
)

func CanEdit(actor string) bool { return actor != "" && actor != "viewer" }
func EnsureActive(r domain.Referee) error {
	if !r.Active {
		return fmt.Errorf("referee disabled")
	}
	return nil
}
func EnsureQualificationOwner(q domain.Qualification, r domain.Referee) error {
	if q.RefereeID != r.ID {
		return fmt.Errorf("owner mismatch")
	}
	return EnsureActive(r)
}
