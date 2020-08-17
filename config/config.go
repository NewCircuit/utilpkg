package config

import (
	"errors"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

// GetConfig will create a YML config if it doesn't exist already.
func GetConfig(path string, output interface{}) error {
	if _, err := os.Stat(path); err != nil {
		return genConfig(path, output)
	}

	// Config file exists, so we're reading it.
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	// Parse the yml file
	_ = yaml.Unmarshal(file, output)

	return nil
}

// Save intakes an interface (which should have YAML meta tags) and saves it to
// the given path.
func Save(path string, output interface{}) error {
	serialized, err := yaml.Marshal(output)

	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, serialized, 0660); err != nil {
		return err
	}

	return nil
}

// genConfig takes <path string> <reference interface{}>
// Generates configuration
// Returns error
func genConfig(path string, reference interface{}) error {
	serialized, err := yaml.Marshal(reference)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, serialized, 0660)

	if err != nil {
		return err
	}

	return errors.New("generated a new default configuration")
}
