package botutil

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
	"regexp"
)

// HasRole Check if a user has a role
func HasRole(has []string, required []string) bool {
	for _, hasRole := range has {
		for _, reqRole := range required {
			if reqRole == hasRole {
				return true
			}
		}
	}
	return false
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

// FilterTag takes <tag string>
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
