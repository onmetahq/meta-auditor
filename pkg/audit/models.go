package auditor

import (
	"strings"
	"time"
)

type Logs struct {
	Path       string    `json:"path"`
	RequestID  string    `json:"requestId"`
	StatusCode int       `json:"statusCode"`
	Request    any       `json:"request"`
	Response   any       `json:"response"`
	CreatedAt  time.Time `json:"created_at"`
	TimeEnded  time.Time `json:"timeEnded"`
	Duration   int64     `json:"duration"`
	Provider   string    `json:"provider"`
}

// MaskString masks a string by replacing characters with asterisks,
func MaskString(s string) string {
	n := len(s)
	if n < 4 {
		return strings.Repeat("*", n)
	}
	return s[:2] + strings.Repeat("*", n-4) + s[n-2:]
}
