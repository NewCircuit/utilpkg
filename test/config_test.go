package utilpkg

import (
	"fmt"
	"strings"
	"testing"

	"github.com/NewCircuit/utilpkg/config"
)

type BotConfig struct {
	Token  string
	Prefix string
}

func TestConfig(t *testing.T) {
	reference := BotConfig{
		"pee pee poo poo",
		".tambourine",
	}
	generic := BotConfig{
		Token:  "",
		Prefix: "",
	}

	err := config.GetConfig("./config.yml", &reference)

	if err != nil {
		if !strings.Contains(err.Error(), "default configuration") {
			t.Error(err)
		} else {
			fmt.Println("Successfully created a default configuration file.")
		}
		return
	}

	err = config.GetConfig("./config.yml", &generic)

	if err != nil {
		t.Error("Failed to read from configuration file.")
		return
	}
	
	if generic.Token != reference.Token {
		fmt.Printf("\"%s\" != \"%s\"\n", generic.Token, reference.Token)
		t.Error("Token attribute does not match reference.")
	} else if generic.Prefix != reference.Prefix {
		fmt.Printf("\"%s\" != \"%s\"\n", generic.Prefix, reference.Prefix)
		t.Error("Prefix attribute does not match reference")
	}
}
