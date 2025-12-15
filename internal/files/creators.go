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

// createMetadataFile creates metadata file (with name) in base
func createMetadataFile(base string, name string) error {
	return os.WriteFile(metadataFile(base, name), []byte{}, 0666)
}

// createNameDir creates name directory (with name) in base
func createNameDir(base string, name string) error {
	return os.Mkdir(nameDir(base, name), 0777)
}

// createDofConfigDir creates config directory in base
func createDofConfigDir(base string) error {
	return os.Mkdir(dofConfigDir(base), 0777)
}

// createDofConfigFile creates config file (with name) in base
func createDofConfigFile(base string, name string) error {
	return os.WriteFile(dofConfigFile(base, name), []byte{}, 0666)
}
