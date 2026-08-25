package domain

import "fmt"

func ValidateReferee(r Referee) error {
	if r.ID == "" {
		return fmt.Errorf("referee id required")
	}
	if r.Name == "" {
		return fmt.Errorf("name required")
	}
	return nil
}
func ValidateQualification(q Qualification) error {
	if q.ID == "" || q.RefereeID == "" {
		return fmt.Errorf("qualification identity required")
	}
	if q.Level == "" || q.Sport == "" {
		return fmt.Errorf("level and sport required")
	}
	if !ValidDate(q.ValidUntil) {
		return fmt.Errorf("invalid expiry")
	}
	return nil
}
func ValidateTraining(t TrainingRecord) error {
	if t.ID == "" || t.RefereeID == "" || t.Course == "" {
		return fmt.Errorf("training fields required")
	}
	if !ValidDate(t.CompletedOn) {
		return fmt.Errorf("invalid training date")
	}
	if t.Score < 0 || t.Score > 100 {
		return fmt.Errorf("score out of range")
	}
	return nil
}
