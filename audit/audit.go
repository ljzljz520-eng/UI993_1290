package audit

import (
	"fmt"
	"refereequal/domain"
	"refereequal/repository"
)

type Logger struct{ R *repository.Repository }

func New(r *repository.Repository) *Logger { return &Logger{R: r} }
func (l *Logger) Log(action, actor, target string) error {
	if action == "" || actor == "" {
		return fmt.Errorf("audit fields required")
	}
	return l.R.Audit(domain.AuditEvent{ID: action + ":" + target, Action: action, Actor: actor, Target: target, At: "2026-01-01"})
}
func (l *Logger) History() ([]domain.AuditEvent, error) { return l.R.S.AllAudits() }
