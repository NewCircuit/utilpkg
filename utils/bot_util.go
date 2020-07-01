package utilitypackage

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

type PromptOptions struct {
	prompt  string
	expires int
}

func PromptUser(s *Session, message *Message, options PromptOptions) (*Message, error) {
	return Prompt(s, message.Author.ID, message.ChannelID, options)
}

func Prompt(s *Session, userID string, channelID string, options PromptOptions) (*Message, error) {
	resChannel := make(chan *Message)
	var response *Message

	_, err := Mention(s, userID, channelID, options.prompt)

	if err != nil {
		return nil, err
	}

	s.AddHandlerOnce(func(s *Session, msg *MessageCreate) {
		go handlePrompt(s, msg, userID, channelID, resChannel)
	})

	response = <-resChannel

	return response, nil
}

func handlePrompt(s *Session, msg *MessageCreate, userID string, channelID string, output chan *Message) {
	if msg.ChannelID == channelID && msg.Author.ID == userID {
		output <- msg.Message
	} else {
		s.AddHandlerOnce(func(s *Session, msg *MessageCreate) {
			handlePrompt(s, msg, userID, channelID, output)
		})
	}
}
