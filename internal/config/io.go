package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// read reads the config file from given filepath
func (c *Config) read(filepath string) error {
	input, err := os.ReadFile(filepath)
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

// write writes to the config file at given filepath
func (c *Config) write(filepath string) error {
	b, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, b, 0666)
	if err != nil {
		return err
	}

	return nil
}
