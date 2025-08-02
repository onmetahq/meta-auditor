package auditor

import "database/sql"

type AuditClient interface {
	AddLogs(logs Logs) error
}

type auditorClient struct {
	client *sql.DB
	table  string
}

func (c *auditorClient) AddLogs(logs Logs) error {
	query := `
		INSERT INTO ` + c.table + ` (path, requestId, statusCode, request, response, created_at, timeEnded, duration, provider)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := c.client.Exec(query, logs.Path, logs.RequestID, logs.StatusCode, logs.Request, logs.Response, logs.CreatedAt, logs.TimeEnded, logs.Duration, logs.Provider)
	return err
}

func NewAuditClient(db *sql.DB, table string) AuditClient {
	return &auditorClient{
		client: db,
		table:  table,
	}
}
