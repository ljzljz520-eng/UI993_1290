package domain

import "strings"

func NormalizeName(v string) string { return strings.TrimSpace(strings.ToLower(v)) }
func SameName(a, b string) bool     { return NormalizeName(a) == NormalizeName(b) }
func Status(active bool) string {
	if active {
		return "active"
	}
	return "disabled"
}
func QualificationKey(q Qualification) string { return q.RefereeID + ":" + q.Sport + ":" + q.Level }
