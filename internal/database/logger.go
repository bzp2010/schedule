package database

import (
	"time"

	"github.com/bzp2010/schedule/internal/log"
	"moul.io/zapgorm2"
)

func newLogger() zapgorm2.Logger {
	logger := zapgorm2.New(log.GetLogger().Desugar())
	logger.SetAsDefault()
	logger.SlowThreshold = time.Second
	return logger
}
