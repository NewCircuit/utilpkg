package logger

// Logger is a type that can be used to abstract away logging and add specific
type Logger interface {
	CreateSubLogger(section string) Logger

	Error(message string)
	Warn(message string)
	Message(message string)
}