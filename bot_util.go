package utility_package

import (
	"fmt"
	. "github.com/bwmarrin/discordgo"
)

func Reply(s *Session, msg *Message, context string) (*Message, error) {
	return s.ChannelMessageSend(
		msg.ChannelID,
		fmt.Sprintf("<@%s> %s", msg.Author.ID, context),
	)
}

func Mention(s *Session, userID string, channelID string, context string) (*Message, error) {
	return s.ChannelMessageSend(
		channelID,
		fmt.Sprintf("<@%s> %s", userID, context),
	)
}
