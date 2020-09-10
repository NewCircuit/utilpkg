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

// Reply sends a response in the same channel to the same user who sent the
// message. The sent message and possible error is returned.
func Reply(s *dg.Session, msg *dg.Message, context string) (*dg.Message, error) {
	return s.ChannelMessageSend(
		msg.ChannelID,
		fmt.Sprintf("<@%s> %s", msg.Author.ID, context),
	)
}

// Mention mentions a given user in a channel.
func Mention(s *dg.Session, userID string, channelID string, context string) (*dg.Message, error) {
	return s.ChannelMessageSend(
		channelID,
		fmt.Sprintf("<@%s> %s", userID, context),
	)
}

// FilterTag takes intakes a ID formatted by the Discord client
// ie <#ID>, <:ID>, or <#ID> and returns ID.
func FilterTag(tag string) string {
	if len(tag) < 2 {
		return tag;
	}
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
