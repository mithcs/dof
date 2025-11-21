package files

import "os"

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

// createNameDir creates name directory in base
func createNameDir(base string, name string) error {
	return os.Mkdir(nameDir(base, name), 0777)
}
