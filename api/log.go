package api

import (
	"fmt"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/hhyouke/server/logger"
	"go.uber.org/zap"
)

type structuredLogger struct {
	Logger *zap.Logger
}

type structuredLoggerEntry struct {
	Logger *zap.Logger
}

func newStructuredLogger(appLogger *logger.AppLogger) func(next http.Handler) http.Handler {
	// chimiddleware.Recoverer()
	return chimiddleware.RequestLogger(&structuredLogger{appLogger.Logger})
}

func (l *structuredLogger) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	entry := &structuredLoggerEntry{Logger: l.Logger}
	// entry.Logger.Core().With()
	entry.Logger = entry.Logger.With(
		zap.String("component", "api"),
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("addr", r.RemoteAddr),
		zap.String("referer", r.Referer()),
	)
	if reqID := GetRequestID(r.Context()); reqID != "" {
		entry.Logger = entry.Logger.With(zap.String("request_id", reqID))
	}
	entry.Logger.Sugar().Info("request started")
	// WithLogEntry(r.Context(), entry)
	return entry
}

func (entry *structuredLoggerEntry) Panic(v interface{}, stack []byte) {
	entry.Logger = entry.Logger.With(
		zap.String("stack", string(stack)),
		zap.String("panic", fmt.Sprintf("%+v", v)),
	)
	entry.Logger.Panic("unhandled request panic")
}

func (entry *structuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	entry.Logger = entry.Logger.With(
		zap.Int("status", status),
		zap.Duration("duration", elapsed),
	)
	entry.Logger.Sugar().Info("request complete")
}

func getLogEntry(r *http.Request) *structuredLoggerEntry {
	entry, _ := chimiddleware.GetLogEntry(r).(*structuredLoggerEntry)
	return entry
}
