package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// CreateLogger will create sugar loggerred zap
func CreateLogger(level string) (*zap.SugaredLogger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stdout"}
	levelAt, ok := logLevel[level]
	// using info level by default
	if !ok {
		levelAt = zap.InfoLevel
	}
	config.Level = zap.NewAtomicLevelAt(levelAt)
	loggerRaw, err := config.Build()
	if err != nil {
		return nil, err
	}
	logger := loggerRaw.Sugar()
	return logger, nil
}

var logLevel = map[string]zapcore.Level{
	"debug":  zap.DebugLevel,
	"info":   zap.InfoLevel,
	"warn":   zap.WarnLevel,
	"error":  zap.ErrorLevel,
	"panic":  zap.PanicLevel,
	"fatal":  zap.FatalLevel,
	"silent": zap.FatalLevel + 1,
}
