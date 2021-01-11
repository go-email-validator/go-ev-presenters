package log

import (
	"go.uber.org/zap"
)

// Default logger logs to console by default
var logger *zap.Logger

func init() {
	builder := zap.NewProductionConfig()
	builder.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	l, err := builder.Build()
	if err != nil {
		panic(err)
	}
	SetLogger(l)
}

func SetLogger(l *zap.Logger) {
	logger = l
}

func Logger() *zap.Logger {
	return logger
}
