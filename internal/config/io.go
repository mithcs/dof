package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// read reads the config file
func read(c *Config, filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// exit if file is empty
	if len(input) == 0 {
		return nil
	}

	err = toml.Unmarshal(input, c)
	if err != nil {
		return err
	}

	return nil
}

// write writes to the config file
func write(c *Config, filename string) error {
	b, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0666)
	if err != nil {
		return err
	}

	return nil
}
