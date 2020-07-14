package decorators

import "github.com/Floor-Gang/utilpkg/logger"

type LogLevel int

const (
	Message = 0
	Warning = 1
	Error = 2
)

type LevelBasedLogFilter struct {
	logger.Logger

	verbosity LogLevel
	child logger.Logger
}

func ApplyLevelBasedLogFilter(verbosity LogLevel, logger logger.Logger) logger.Logger {
	decoratedLogger := LevelBasedLogFilter{
		verbosity: verbosity,
		child: logger,
	}
	return decoratedLogger.Logger
}

// CreateSubLogger created a logger that references where the initial logger was created.
func (logger *LevelBasedLogFilter) CreateSubLogger(section string) logger.Logger {
	return logger.CreateSubLogger(section)
}

func (logger *LevelBasedLogFilter) Warn(message string) {
	if logger.verbosity <= Warning {
		logger.Warn(message)
	}
}

func (logger *LevelBasedLogFilter) Error(message string) {
	if logger.verbosity <= Error {
		logger.Warn(message)
	}
}

func (logger *LevelBasedLogFilter) Message(message string) {
	if logger.verbosity <= Message {
		logger.Warn(message)
	}
}
