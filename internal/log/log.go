package log

import (
	"os"

	"github.com/bzp2010/schedule/internal/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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
	var (
		logPath = config.Log.Path
		level   zap.AtomicLevel
		err     error
	)

	// convert log level to zap format
	if level, err = zap.ParseAtomicLevel(config.Log.Level); err != nil {
		return errors.Wrap(err, "failed to parse log level")
	}

	encConfig := zap.NewProductionEncoderConfig()
	coreFile := zapcore.NewCore(
		zapcore.NewJSONEncoder(encConfig),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    10, // maximum size of a single log file (Mbytes)
			MaxBackups: 10, // maximum number of logs to be saved
			MaxAge:     30, // maximum number of days to keep logs
			Compress:   false,
		}),
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return level.Enabled(l)
		}),
	)

	var core zapcore.Core
	if config.Debug {
		// print logs to stdout when debug mode is on
		coreConsole := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encConfig),
			zapcore.AddSync(os.Stdout),
			zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				return level.Enabled(l)
			}),
		)
		core = zapcore.NewTee(coreFile, coreConsole)
	} else {
		core = zapcore.NewTee(coreFile)
	}

	// build logger
	zapLog := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	defer zapLog.Sync()

	logger = zapLog.Sugar()
	return nil
}
