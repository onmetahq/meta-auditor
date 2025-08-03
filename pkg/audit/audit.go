package auditor

import (
	"database/sql"
	"strings"
)

type AuditClient interface {
	AddLogs(logs Logs) error
}

type auditorClient struct {
	client  *sql.DB
	table   string
	service string
}

func (c *auditorClient) AddLogs(logs Logs) error {
	query := `
		INSERT INTO ` + c.table + ` (path, requestId, statusCode, request, response, created_at, timeEnded, duration, provider, service)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := c.client.Exec(query, logs.Path, logs.RequestID, logs.StatusCode, logs.Request, logs.Response, logs.CreatedAt, logs.TimeEnded, logs.Duration, logs.Provider, c.service)
	return err
}

func NewAuditClient(db *sql.DB, table, service string) AuditClient {
	return &auditorClient{
		client:  db,
		table:   table,
		service: service,
	}
}

func MaskString(s string) string {
	n := len(s)
	if n < 4 {
		return strings.Repeat("*", n)
	}
	return s[:2] + strings.Repeat("*", n-4) + s[n-2:]
}
