package utilpkg

import (
	"fmt"
)

// Reply takes <s *Session> <msg *Message> <context string>
// Sends message to channel
// Returns replied message and error
func Reply(s *Session, msg *Message, context string) (*Message, error) {
	return s.ChannelMessageSend(
		msg.ChannelID,
		fmt.Sprintf("<@%s> %s", msg.Author.ID, context),
	)
}

// Mention takes <s *Session> <userID string> <channelID string> <context string>
// Mentions user in channel
// Returns replied message and error
func Mention(s *Session, userID string, channelID string, context string) (*Message, error) {
	return s.ChannelMessageSend(
		channelID,
		fmt.Sprintf("<@%s> %s", userID, context),
	)
}

// PromptOptions structure
type PromptOptions struct {
	prompt  string
	expires int
}

// PromptUser takes <s *Session> <message *Message> <options PromptOptions>
// Prompts user
// Returns message and error
func PromptUser(s *Session, message *Message, options PromptOptions) (*Message, error) {
	return Prompt(s, message.Author.ID, message.ChannelID, options)
}

// Prompt takes <s *Session> <userID string> <channelID string> <options PromptOptions>
// Prompts user for response
// Returns message and error
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

// handlePrompt takes <s *Session> <msg *MessageCreate> <userID string> <channelID string> <output chan *message>
// Handles prompt by adding handler
func handlePrompt(s *Session, msg *MessageCreate, userID string, channelID string, output chan *Message) {
	if msg.ChannelID == channelID && msg.Author.ID == userID {
		output <- msg.Message
	} else {
		s.AddHandlerOnce(func(s *Session, msg *MessageCreate) {
			handlePrompt(s, msg, userID, channelID, output)
		})
	}
}
