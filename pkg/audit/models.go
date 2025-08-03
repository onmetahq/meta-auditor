package auditor

import (
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
