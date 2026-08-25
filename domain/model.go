package domain

import "time"

type Referee struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at"`
}
type Qualification struct {
	ID         string `json:"id"`
	RefereeID  string `json:"referee_id"`
	Level      string `json:"level"`
	Sport      string `json:"sport"`
	ValidUntil string `json:"valid_until"`
	Issuer     string `json:"issuer"`
	Active     bool   `json:"active"`
}
type TrainingRecord struct {
	ID          string `json:"id"`
	RefereeID   string `json:"referee_id"`
	Course      string `json:"course"`
	CompletedOn string `json:"completed_on"`
	Score       int    `json:"score"`
}
type QueryResult struct {
	Referee        Referee          `json:"referee"`
	Qualifications []Qualification  `json:"qualifications"`
	Trainings      []TrainingRecord `json:"trainings"`
}
type AuditEvent struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Actor  string `json:"actor"`
	Target string `json:"target"`
	At     string `json:"at"`
}

func ParseDate(v string) (time.Time, error)    { return time.Parse("2006-01-02", v) }
func ValidDate(v string) bool                  { _, e := ParseDate(v); return e == nil }
func (q Qualification) IsValid(on string) bool { return q.Active && q.ValidUntil >= on }
