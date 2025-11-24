package files

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CreateDotfilesDir creates dotfiles directory
func CreateDotfilesDir() error {
	base := dataDir()
	return createDotfilesDir(base)
}

// CreateDofDir creates .dof directory
func CreateDofDir() error {
	base := dataDir()
	return createDofDir(base)
}

// CreateMetadataFile creates metadata file
func CreateMetadataFile(name string) error {
	base := dataDir()
	return createMetadataFile(base, name)
}

// MetadataPath returns path to metadata file
func MetadataPath(name string) string {
	base := dataDir()
	return metadataFile(base, name)
}

// GeneralizePaths returns generalized paths, which replaces machine (and OS) specific
// part of paths with machine (and OS) independent values
func GeneralizePaths(paths []string) []string {
	config := configDir()
	profile := profileDir()

	for i, path := range paths {
		path = strings.Replace(path, config, "{CONFIG}", 1)
		path = strings.Replace(path, profile, "{PROFILE}", 1)

		path = strings.ReplaceAll(path, string(filepath.Separator), "{.}")
		paths[i] = path
	}

	return paths
}

// CopyToName copies files to dotfiles directory, starting file names from <start>
func CopyToName(paths []string, name string, start int) error {
	base := dataDir()
	err := createNameDir(base, name)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	for _, path := range paths {
		dest := filepath.Join(dotfilesDir(base), name, strconv.Itoa(start))
		err = copyFileTree(path, dest)
		if err != nil {
			return err
		}

		start += 1
	}

	return nil
}

// MoveAndSymlink moves to dotfiles directory and symlinks to source
// starting file names from <start>
func MoveAndSymlink(paths []string, name string, start int) error {
	base := dataDir()
	err := createNameDir(base, name)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}

	for _, path := range paths {
		dest := filepath.Join(dotfilesDir(base), name, strconv.Itoa(start))
		err = replaceWithSymlink(path, dest)
		if err != nil {
			return err
		}

		start += 1
	}
	return nil
}
