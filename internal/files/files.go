package files

import (
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

// GeneralizePaths returns generalized paths
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
	createNameDir(base, name)

	for _, path := range paths {
		dest := filepath.Join(dotfilesDir(base), name, strconv.Itoa(start))
		if err := copyFileTree(path, dest); err != nil {
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

	for _, path := range paths {
		dest := filepath.Join(dotfilesDir(base), name, strconv.Itoa(start))
		replaceWithSymlink(path, dest)
		start += 1
	}
	return nil
}
