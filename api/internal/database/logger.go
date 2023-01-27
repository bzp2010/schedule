package database

import (
	"time"

	"github.com/bzp2010/schedule/internal/config"
	"github.com/bzp2010/schedule/internal/log"
	"moul.io/zapgorm2"
)

func newLogger(cfg config.Config) zapgorm2.Logger {
	logger := zapgorm2.New(log.GetLogger().Desugar())
	logger.SetAsDefault()

	if cfg.Debug {
		// when debug mode is on, all SQL queries will be logged
		logger.SlowThreshold = time.Nanosecond
	}

	return logger
}
