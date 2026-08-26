package fixtures

import "refereequal/domain"

func Referee() domain.Referee {
	return domain.Referee{ID: "R-100", Name: "Lin Mei", Active: true, CreatedAt: "2026-01-01"}
}
func Qualification(id, level, sport string) domain.Qualification {
	return domain.Qualification{ID: id, RefereeID: "R-100", Level: level, Sport: sport, ValidUntil: "2027-12-31", Issuer: "National Board", Active: true}
}
func Training(id string, score int) domain.TrainingRecord {
	return domain.TrainingRecord{ID: id, RefereeID: "R-100", Course: "Rules Refresher", CompletedOn: "2026-02-01", Score: score}
}
