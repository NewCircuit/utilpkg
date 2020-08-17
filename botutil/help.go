package botutil

import (
	"fmt"
	dg "github.com/bwmarrin/discordgo"
)

// Feature represents your bot.
type Feature struct {
	// Name of your feature ie. "Nickname Manager"
	Name string
	// Description for describing your feature "Moderates nicknames"
	Description string
	// Commands are all the possible commands that exist, optionally empty.
	Commands []*Command
	// Prefix is the prefix for your bot.
	Prefix string
}

// Command are all the possible commands that exist.
type Command struct {
	// Name is the command name that is callable by a Discord user
	// ie. "request"
	Name string
	// Description describes your command's usage
	// ie. "This is for request a new nickname"
	Description string
	// Example is example usage of the command
	// ie. [".nick", "request", "new nickname"]
	Example []string
}

// BuildHelp makes a help message embed.
func BuildHelp(feature Feature) (embed *dg.MessageEmbed) {
	embed = &dg.MessageEmbed{
		Title:       feature.Name,
		Description: feature.Description,
		Color:       0xef2f2f,
		Fields:      []*dg.MessageEmbedField{},
	}

	for _, command := range feature.Commands {
		field := &dg.MessageEmbedField{
			Name: command.Name,
		}
		field.Value += command.Description + "\n"

		length := len(command.Example)
		for i, example := range command.Example {
			if i == 0 {
				field.Value += fmt.Sprintf("%s", example)
			} else {
				field.Value += fmt.Sprintf("<%s>", example)
			}

			if (length - 1) != i {
				field.Value += " "
			}
		}
		embed.Fields = append(embed.Fields, field)
	}

	return embed
}
