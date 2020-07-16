package logger

// Logger is a type that can be used to abstract away logging and add specific
type Logger interface {
	// CreateSubLogger creates a nested instance of a logger that tracks its state.
	CreateSubLogger(section string) Logger

	// Debug logs a debug message.
	Debug(message string)

	// Error logs an error message.
	Error(error error)

	// Warn logs a warning.
	Warn(message string)

	// Message logs a message.
	Message(message string)
}