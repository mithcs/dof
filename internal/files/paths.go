package files

import "path/filepath"

// dotfilesDir returns path to dotfiles dir in base
func dotfilesDir(base string) string {
	return filepath.Join(base, "dotfiles")
}

// dofDir returns path to .dof dir in base
func dofDir(base string) string {
	dotfiles := dotfilesDir(base)
	return filepath.Join(dotfiles, ".dof")
}

// metadataFile returns path to metadata file in base
func metadataFile(base string, name string) string {
	dof := dofDir(base)
	return filepath.Join(dof, name)
}
