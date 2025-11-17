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
