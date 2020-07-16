package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// StreamLogger logs to streams (e.g. console screen, network connections, files).
type StreamLogger struct {
	Logger

	section string

	out *io.Writer
	err *io.Writer
}

// NewConsoleLoggerBasic is a factory method to create a logger that outputs to the console screen.
func NewConsoleLoggerBasic(section string) Logger {
	outputWriter := io.Writer(os.Stdout)
	errorWriter := io.Writer(os.Stderr)
	return NewConsoleLogger(section, &outputWriter, &errorWriter)
}

// NewConsoleLogger is a factory method to set up a console logger that explicitly outputs to the OS determined output
// streams.
func NewConsoleLogger(section string, out *io.Writer, err *io.Writer) Logger {
	logger := StreamLogger{
		section: section,
		out:     out,
		err:     err,
	}
	return logger.Logger
}

func (logger *StreamLogger) CreateSubLogger(section string) Logger {
	return NewConsoleLogger(logger.section+":"+section, logger.out, logger.err)
}

func (logger *StreamLogger) Debug(message string) {
	write(message, "debug", *logger.out)
}

func (logger *StreamLogger) Error(error error) {
	write(error.Error(), "error", *logger.err)
}

func (logger *StreamLogger) Message(message string) {
	write(message, "message", *logger.out)
}

func (logger *StreamLogger) Warn(message string) {
	write(message, "warning", *logger.out)
}

func write(message string, level string, dest io.Writer) {
	formattedMessage := fmt.Sprintf("%s: %s", level, message)
	_, err := dest.Write([]byte(formattedMessage))
	if err != nil {
		log.Panic("Failed to write to stream.")
	}
}
