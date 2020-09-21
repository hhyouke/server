package api

import (
	"context"

	"github.com/hhyouke/server/conf"
)

type contextKey string

func (c contextKey) String() string {
	return "api context key " + string(c)
}

const (
	tokenKey     = contextKey("jwt")
	configKey    = contextKey("configuration")
	requestIDKey = contextKey("request_id")
	userIDKey    = contextKey("user_id")
	userKey      = contextKey("user")
	dbKey        = contextKey("db")
	loggerKey    = contextKey("logger")
	machineKey   = contextKey("machine")
)

// // WithLogger adds logger to the context
// func WithLogger(ctx context.Context, logger *logger.AppLogger) context.Context {
// 	return context.WithValue(ctx, loggerKey, logger)
// }

// // GetLogger get logger from context
// func GetLogger(ctx context.Context) *logger.AppLogger {
// 	obj := ctx.Value(loggerKey)
// 	if obj == nil {
// 		return nil
// 	}
// 	return obj.(*logger.AppLogger)
// }

// WithConfig adds configuration to the context.
func WithConfig(ctx context.Context, conf *conf.GlobalConfiguration) context.Context {
	return context.WithValue(ctx, configKey, conf)
}

// GetConfig reads the global configuration from the context.
func GetConfig(ctx context.Context) *conf.GlobalConfiguration {
	obj := ctx.Value(configKey)
	if obj == nil {
		return nil
	}

	return obj.(*conf.GlobalConfiguration)
}

// WithRequestID adds the provided request ID to the context.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// GetRequestID reads the request ID from the context.
func GetRequestID(ctx context.Context) string {
	obj := ctx.Value(requestIDKey)
	if obj == nil {
		return ""
	}

	return obj.(string)
}

// WithMachineID adds the machine ID to the context.
func WithMachineID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, machineKey, id)
}

// GetMachineID reads the machine ID from the context.
func GetMachineID(ctx context.Context) string {
	obj := ctx.Value(machineKey)
	if obj == nil {
		return ""
	}

	return obj.(string)
}
