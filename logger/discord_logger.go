package logger

import (
	"errors"
	"github.com/Floor-Gang/utilpkg/botutil"
	"log"
)

type DiscordLogger struct {
	Logger

	baseEmbedBuilder func() *botutil.Embed

	section string

	useWebhook bool
	webhookURL string
}

func NewDiscordLoggerFromWebhook(section string, webhookURL string) DiscordLogger {
	return DiscordLogger{
		baseEmbedBuilder: func() *botutil.Embed {
			return botutil.NewEmbed()
		},
		section: section,
		webhookURL: webhookURL,
		useWebhook: true,
	}
}
func NewDiscordLoggerFromWebhookWithEmbed(section string, webhookURL string, baseEmbed func() *botutil.Embed) DiscordLogger {
	return DiscordLogger{
		baseEmbedBuilder: baseEmbed,
		section: section,
		webhookURL: webhookURL,
		useWebhook: true,
	}
}

func (logger *DiscordLogger) CreateSubLogger(section string) DiscordLogger {
	if logger.useWebhook {
		return NewDiscordLoggerFromWebhookWithEmbed(
			logger.section + " -> " + section,
			logger.webhookURL,
			logger.baseEmbedBuilder)
	}
	// TODO(velddev): Create channel-based logger with token instead of webhook URL.
	return DiscordLogger{}
}

func (logger *DiscordLogger) Warn(message string) {
	_ = logger.sendWebhook(message, "warning")
}

func (logger *DiscordLogger) Error(message string) {
	_ = logger.sendWebhook(message, "error")
}

func (logger *DiscordLogger) Message(message string) {
	_ = logger.sendWebhook(message, "message")
}

func (logger *DiscordLogger) sendWebhook(message string, level string) error {
	if !logger.useWebhook {
		return errors.New("tried to push webhook, but logger was not setup with webhook")
	}

	embed := logger.baseEmbedBuilder()

	embed.SetAuthor(level, "", "")
	embed.SetTitle(mergeStrings(logger.section, embed.Title))
	embed.SetDescription(mergeStrings(message, embed.Description))
	err := embed.SetColor(determineColorFromLevel(level))
	if err != nil {
		log.Panic(err)
	}
	return embed.SendToWebhook(logger.webhookURL)
}

func determineColorFromLevel(level string) string {
	switch level {
		case "error": return "15086631"
		case "warning": return "15459356"
		case "message": return "1875179"
		default: return "16777215"
	}
}

func mergeStrings(a string, b string) string {
	if a == "" {
		return b
	}

	if b == "" {
		return a
	}

	return a + " (" + b + ")"
}
