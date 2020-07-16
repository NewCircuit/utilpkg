package decorators

import "github.com/Floor-Gang/utilpkg/logger"

type LogLevel int

const (
	Debug = -1
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
		logger.child.Warn(message)
	}
}

func (logger *LevelBasedLogFilter) Error(error error) {
	if logger.verbosity <= Error {
		logger.child.Error(error)
	}
}

func (logger *LevelBasedLogFilter) Message(message string) {
	if logger.verbosity <= Message {
		logger.child.Message(message)
	}
}

func (logger *LevelBasedLogFilter) Debug(message string) {
	if logger.verbosity <= Debug {
		logger.child.Debug(message)
	}
}
