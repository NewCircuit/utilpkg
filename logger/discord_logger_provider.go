package logger

type DiscordProvider struct {
	Provider
}

func (provider *DiscordProvider) NewLogger(section string) Logger {
	return NewConsoleLoggerBasic(section)
}