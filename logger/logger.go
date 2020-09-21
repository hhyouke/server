package logger

import (
	"log"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLevel = "info"

// AppLogger the basic logger
type AppLogger struct {
	*zap.Logger
}

// NewLogger create a new logger
func NewLogger(level string) *AppLogger {
	logLevel := getLogLevel(level)
	writeSyncer := getLogWriter()
	consoleWriter := zapcore.Lock(os.Stdout)
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, logLevel),
		zapcore.NewCore(encoder, consoleWriter, logLevel),
	)
	logger := zap.New(core, zap.AddCaller())
	appLogger := &AppLogger{
		logger,
	}
	return appLogger
}

func getLogLevel(level string) zapcore.Level {
	var lvl string
	if level == "" {
		log.Println("no level input, switch to default level(info)")
		lvl = defaultLevel
	} else {
		lvl = level
	}
	switch {
	case lvl == "debug":
		return zapcore.DebugLevel
	case lvl == "info":
		return zapcore.InfoLevel
	case lvl == "warn":
		return zapcore.WarnLevel
	case lvl == "error":
		return zapcore.ErrorLevel
	case lvl == "panic":
		return zapcore.PanicLevel
	}
	return 0
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    20,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
