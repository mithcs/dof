package config

import "github.com/mithcs/dof/internal/files"

type Config struct {
	Method string `toml:"method"`
}

var filePath string = files.ConfigFile("config.toml")

// Create creates configuration file
func (c *Config) Create() error {
	// name of config dir
	dir := "dof"

	err := files.CreateConfigDirectory(dir)
	if err != nil {
		return err
	}

	err = files.CreateConfigFile(filePath)
	if err != nil {
		return err
	}

	c.setDefaults()

	err = c.write(filePath)
	if err != nil {
		return err
	}

	return nil
}

// DefaultMethod returns method from config
func (c *Config) DefaultMethod() (string, error) {
	err := c.read(filePath)
	if err != nil {
		return "", err
	}

	return c.Method, nil
}

// setDefaults sets default config values
func (c *Config) setDefaults() {
	c.Method = "copy"
}
