// Package logger provides helper functions for using zap.Logger.
package logger

import (
	"fmt"
	"log"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const logLevelEnvName = "LOG_LEVEL"

//nolint:gochecknoglobals
var (
	errr    error
	loggerr *zap.Logger
	once    sync.Once
)

// I is the alias for more convenient using.
var I = MustGetLogger

// GetLogger is the provider for zap.Logger.
// This function init logger on the first call.
func GetLogger() (*zap.Logger, error) {
	once.Do(func() {
		loggerr, errr = createLogger(os.Getenv(logLevelEnvName))
	})

	return loggerr, errr
}

// MustGetLogger is the provider for zap.Logger.
// This function init logger on the first call.
// It's function occurs the panic if zap.Logger can not to initialize.
func MustGetLogger() *zap.Logger {
	logger, err := GetLogger()
	if err != nil {
		panic(err)
	}

	return logger
}

// Close the zap.Logger.
// You must call this function on shutdown.
func Close() {
	if err := loggerr.Sync(); err != nil {
		log.Printf("can not to flush logger: %s", err.Error())
	}
}

func createLogger(logLevel string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.NameKey = "name"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if logLevel != "" {
		expectedLevel, err := zapcore.ParseLevel(logLevel)
		if err != nil {
			return nil, fmt.Errorf(
				"can not to parse LOG_LEVEL %q, error: %w",
				logLevel, err,
			)
		}

		cfg.Level = zap.NewAtomicLevelAt(expectedLevel)
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf(
			"can not to build config for zap logger, error: %w", err,
		)
	}

	return logger, nil
}
