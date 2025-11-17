package metadata

import (
	"github.com/mithcs/dof/internal/files"
)

type method int

type Entry struct {
	Name   string   `json:"name"`
	Paths  []string `json:"paths"`
	Method method   `json:"method"`
}

type Metadata struct {
	Entries []Entry `json:"entries"`
}

const (
	Copy method = iota
	Symlink
)

var filename string = "metadata.json"

// Create creates the metadata file
func (m *Metadata) Create() error {
	err := files.CreateMetadataFile(filename)
	if err != nil {
		return err
	}

	return nil
}

// Add adds entry to metadata
func (m *Metadata) Add(e Entry) error {
	err := read(m, files.MetadataPath(filename))
	if err != nil {
		return err
	}

	err = add(m, e)
	if err != nil {
		return err
	}

	err = write(m, files.MetadataPath(filename))
	if err != nil {
		return err
	}

	return nil
}

// Remove removes entry from metadata
func (m *Metadata) Remove(name string) error {
	err := read(m, files.MetadataPath(filename))
	if err != nil {
		return err
	}

	err = remove(m, name)
	if err != nil {
		return err
	}

	err = write(m, files.MetadataPath(filename))
	if err != nil {
		return err
	}

	return nil
}
