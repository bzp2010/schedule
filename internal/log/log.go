package log

import (
	"github.com/bzp2010/schedule/internal/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

// GetLogger to get logger instance and creates a new instance when it is not initialized
func GetLogger() *zap.SugaredLogger {
	return logger
}

// SetupLogger will create a single instance of logger based on the configuration
func SetupLogger(config config.Config) error {
	zapLog, err := zap.NewProduction()
	if err != nil {
		return errors.Wrap(err, "failed to new zap logger")
	}
	defer zapLog.Sync()

	logger = zapLog.Sugar()
	return nil
}
