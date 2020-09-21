package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

// DBLogger logger for database operations
type DBLogger struct {
	*zap.Logger
}

// NewDBLogger create new database logger
func NewDBLogger(log *zap.Logger) *DBLogger {
	return &DBLogger{
		log,
	}
}

// Print method of the DBLogger
func (dbl *DBLogger) Print(params ...interface{}) {
	if len(params) <= 1 {
		return
	}

	level := params[0]
	log := dbl.With(zap.Any("gorm_level", level), zap.Any("db_src", params[1]))

	if level != "sql" {
		log.Sugar().Debug(params[2:]...)
		return
	}

	dur := params[2].(time.Duration)
	sql := params[3].(string)
	sqlValues := params[4].([]interface{})
	rows := params[5].(int64)

	values := ""
	if valuesJSON, err := json.Marshal(sqlValues); err == nil {
		values = string(valuesJSON)
	} else {
		values = fmt.Sprintf("%+v", sqlValues)
	}

	log = log.With(
		zap.Int64("dur_ns", dur.Nanoseconds()),
		zap.Duration("dur", dur),
		zap.String("sql", strings.ReplaceAll(sql, `"`, `'`)),
		zap.String("sql", strings.ReplaceAll(values, `"`, `'`)),
		zap.Int64("rows", rows),
	)
	log.Debug("sql query")
}
