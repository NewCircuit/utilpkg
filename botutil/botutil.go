// This package evolves around discordgo library
package botutil

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	"regexp"
)

// Check if a user has a role
func HasRole(has []string, required []string) (bool, string) {
	for _, hasRole := range has {
		for _, reqRole := range required {
			if reqRole == hasRole {
				return true, ""
			}
		}
	}
	return false, ""
}

// Reply takes <s *Session> <msg *Message> <context string>
// Sends message to channel
// Returns replied message and error
func Reply(s *dg.Session, msg *dg.Message, context string) (*dg.Message, error) {
	return s.ChannelMessageSend(
		msg.ChannelID,
		fmt.Sprintf("<@%s> %s", msg.Author.ID, context),
	)
}

// Mention takes <s *Session> <userID string> <channelID string> <context string>
// Mentions user in channel
// Returns replied message and error
func Mention(s *dg.Session, userID string, channelID string, context string) (*dg.Message, error) {
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
func PromptUser(s *dg.Session, message *dg.Message, options PromptOptions) (*dg.Message, error) {
	return Prompt(s, message.Author.ID, message.ChannelID, options)
}

// Prompt takes <s *Session> <userID string> <channelID string> <options PromptOptions>
// Prompts user for response
// Returns message and error
func Prompt(s *dg.Session, userID string, channelID string, options PromptOptions) (*dg.Message, error) {
	resChannel := make(chan *dg.Message)
	var response *dg.Message

	_, err := Mention(s, userID, channelID, options.prompt)

	if err != nil {
		return nil, err
	}

	s.AddHandlerOnce(func(s *dg.Session, msg *dg.MessageCreate) {
		go handlePrompt(s, msg, userID, channelID, resChannel)
	})

	response = <-resChannel

	return response, nil
}

// handlePrompt takes <s *Session> <msg *MessageCreate> <userID string> <channelID string> <output chan *message>
// Handles prompt by adding handler
func handlePrompt(s *dg.Session, msg *dg.MessageCreate, userID string, channelID string, output chan *dg.Message) {
	if msg.ChannelID == channelID && msg.Author.ID == userID {
		output <- msg.Message
	} else {
		s.AddHandlerOnce(func(s *dg.Session, msg *dg.MessageCreate) {
			handlePrompt(s, msg, userID, channelID, output)
		})
	}
}

// filterTag takes <tag string>
// filterTag filters out the random characters
// Returns id
func FilterTag(tag string) string {
	typeTag := tag[1:2]
	m := regexp.MustCompile("")

	switch typeTag {
	// A channel mention
	case "#":
		m = regexp.MustCompile("#(.*?)>")
		break
	// A emoij
	case ":":
		m = regexp.MustCompile(":(.*?)>")
		break
	// A role or a user mention
	case "@":
		if tag[2:3] == "&" {
			m = regexp.MustCompile("&(.*?)>")
		} else {
			m = regexp.MustCompile("!(.*?)>")
		}
		break
	default:
		return tag
		break
	}

	tag = m.FindString(tag)

	// Remove the first and last character
	sz := len(tag)

	if sz > 0 {
		// Remove the last character, which is ">"
		tag = tag[:sz-1]

		// If it's not an emoij, because we need the :emoij:<id>
		if typeTag != ":" {
			// Remove the first character, which is ["!", "&", "#"]
			tag = tag[1:]
		}
	}

	return tag
}
