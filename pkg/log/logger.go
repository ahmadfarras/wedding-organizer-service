package log

import (
	"fartech/wedding-organizer-service/pkg/model"
	"log"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLog(config *model.Config) (logger *zap.Logger) {
	zapConfig := zap.Config{
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	if strings.EqualFold(config.Environment, "production") {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	} else {
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	}

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}

	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	return logger
}
