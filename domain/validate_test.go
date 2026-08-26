package domain

import "testing"

func TestValidation(t *testing.T) {
	if ValidateReferee(Referee{}) == nil {
		t.Fatal()
	}
	if ValidateQualification(Qualification{}) == nil {
		t.Fatal()
	}
	if ValidateTraining(TrainingRecord{}) == nil {
		t.Fatal()
	}
}
