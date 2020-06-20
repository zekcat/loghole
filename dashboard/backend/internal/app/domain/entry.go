package domain

import (
	"encoding/json"
	"time"
)

type Entry struct {
	Time        time.Time       `json:"time"`
	NSec        int64           `json:"nsec"`
	Namespace   string          `json:"namespace"`
	Source      string          `json:"source"`
	Host        string          `json:"host"`
	TraceID     string          `json:"trace_id"`
	Message     string          `json:"message"`
	Params      json.RawMessage `json:"params"`
	BuildCommit string          `json:"build_commit"`
	ConfigHash  string          `json:"config_hash"`
}
