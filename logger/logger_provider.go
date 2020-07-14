package logger

type Provider interface {
	NewLogger(section string) *Logger
}