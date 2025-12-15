package config

import "github.com/mithcs/dof/internal/files"

type Config struct {
	Method string `toml:"method"`
}

var fileName string = "config.toml"
var filePath string = files.ConfigFile(fileName)

// Create creates configuration file with default values
func (c *Config) Create() error {
	dir := "dof"
	err := files.CreateConfigDirectory(dir)
	if err != nil {
		return err
	}

	err = files.CreateConfigFile(fileName)
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
