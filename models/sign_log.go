package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// SignLogs signlog data
type SignLogs struct {
	Log   string `sql:"-"`
	Count uint
	Logs  []SignLog
}

// SignLog signlogs item
type SignLog struct {
	Agent  string
	At     *time.Time
	IP     string
	UID    string
	OrgUID string
	ID     int64
}

// Scan scan data into sign logs
func (signLogs *SignLogs) Scan(data interface{}) (err error) {
	switch values := data.(type) {
	case []byte:
		if string(values) != "" {
			return json.Unmarshal(values, signLogs)
		}
	case string:
		return signLogs.Scan([]byte(values))
	case []string:
		for _, str := range values {
			if err := signLogs.Scan(str); err != nil {
				return err
			}
		}
	default:
		err = errors.New("unsupported driver -> Scan pair for SignLogs")
	}
	return
}

// Value return struct's Value
func (signLogs SignLogs) Value() (driver.Value, error) {
	results, err := json.Marshal(signLogs)
	return string(results), err
}
