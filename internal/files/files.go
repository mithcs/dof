package files

import (
	"os"
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

// createDotfilesDir creates dotfiles directory in base
func createDotfilesDir(base string) error {
	return os.Mkdir(dotfilesDir(base), 0777)
}

// createDofDir creates .dof directory in base
func createDofDir(base string) error {
	return os.Mkdir(dofDir(base), 0777)
}

// createMetadataFile creates metadata file in base
func createMetadataFile(base string, name string) error {
	return os.WriteFile(metadataFile(base, name), []byte{}, 0666)
}
