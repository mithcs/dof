//go:build linux

package files

import (
	"os"
	"path/filepath"
)

// dataDir returns path to store app data to
func dataDir() string {
	dir := os.Getenv("XDG_DATA_HOME")
	if dir == "" {
		home := os.Getenv("HOME")
		dir = filepath.Join(home, ".local", "share")
	}

	return dir
}

// configDir returns path to config directory
func configDir() string {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home := os.Getenv("HOME")
		dir = filepath.Join(home, ".config")
	}

	return dir
}

// profileDir returns path to config directory
func profileDir() string {
	return os.Getenv("HOME")
}
