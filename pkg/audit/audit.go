package auditor

import (
	"database/sql"
	"encoding/json"
	"strings"

	http_models "github.com/onmetahq/meta-http/pkg/models"
)

type AuditClient interface {
	AddLogs(logs Logs) error
}

type auditorClient struct {
	client  *sql.DB
	table   string
	service string
}

func toJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func CastError(err error) map[string]any {
	if err == nil {
		return map[string]any{}
	}

	errInfo, ok := err.(*http_models.HttpClientErrorResponse)
	if !ok {
		newErr := map[string]any{"error": err.Error()}
		return newErr
	}
	newErr := make(map[string]any)
	err = json.Unmarshal([]byte(errInfo.Err.Message), &newErr)
	if err != nil {
		newErr := map[string]any{"error": errInfo.Err.Message}
		return newErr
	}
	return newErr
}

func (c *auditorClient) AddLogs(logs Logs) error {
	query := `
		INSERT INTO ` + c.table + ` (path, request_id, status_code, request, response, created_at, duration_ms, provider, service)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := c.client.Exec(query, logs.Path, logs.RequestID, logs.StatusCode, toJSON(logs.Request), toJSON(logs.Response), logs.CreatedAt, logs.Duration, logs.Provider, c.service)
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
