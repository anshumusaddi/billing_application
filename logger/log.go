package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

var log *zap.SugaredLogger

func InitLogger(logLevel string) {
	if len(logLevel) == 0 {
		logLevel = "info"
	}
	logger, _ := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(getLogLevel(logLevel)),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{TimeKey: "timestamp", EncodeTime: zapcore.ISO8601TimeEncoder,
			MessageKey: "message", LevelKey: "level", EncodeLevel: zapcore.LowercaseLevelEncoder},
	}.Build()
	log = logger.Sugar()
}

func getLogLevel(logLevel string) zapcore.Level {
	switch strings.ToLower(logLevel) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func Debug(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	log.Fatalf(msg, args...)
}

func Info(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Warnf(msg, args...)
}

func Sync() {
	err := log.Sync()
	if err != nil {
		return
	}
}
