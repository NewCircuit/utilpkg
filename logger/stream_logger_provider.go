package logger

type StreamProvider struct {
	Provider
}

func (provider *StreamProvider) NewLogger(section string) Logger {
	return NewConsoleLoggerBasic(section)
}