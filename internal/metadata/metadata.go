package metadata

import (
	"github.com/mithcs/dof/internal/files"
)

// Method refers to method to use - Copy or Symlink
type Method string

// Entry contains group of paths, method and a name for entry
type Entry struct {
	Name   string   `json:"name"`
	Paths  []string `json:"paths"`
	Method Method   `json:"method"`
}

// Metadata is layout of metadata file
type Metadata struct {
	Entries []Entry `json:"entries"`
}

const (
	Copy    Method = "copy"
	Symlink Method = "symlink"
)

var fileName string = "metadata.json"
var filePath string = files.MetadataPath(fileName)

// Create creates the metadata file
func (m *Metadata) Create() error {
	err := files.CreateMetadataFile(fileName)
	if err != nil {
		return err
	}

	return nil
}

// Add adds given entry to metadata
func (m *Metadata) Add(e Entry) error {
	err := m.read(filePath)
	if err != nil {
		return err
	}

	err = m.add(e)
	if err != nil {
		return err
	}

	err = m.write(filePath)
	if err != nil {
		return err
	}

	return nil
}

// Remove removes entry (with given name) from metadata
func (m *Metadata) Remove(name string) error {
	err := m.read(filePath)
	if err != nil {
		return err
	}

	err = m.remove(name)
	if err != nil {
		return err
	}

	err = m.write(filePath)
	if err != nil {
		return err
	}

	return nil
}

// Get returns entry (with given name) from metadata
func (m *Metadata) Get(name string) (Entry, error) {
	err := m.read(filePath)
	if err != nil {
		return Entry{}, err
	}

	return m.get(name)
}

// Update replaces existing entry with updated entry in metadata
func (m *Metadata) Update(updated Entry) error {
	if len(updated.Paths) < 1 {
		return m.Remove(updated.Name)
	}

	err := m.read(filePath)
	if err != nil {
		return err
	}

	err = m.update(updated)
	if err != nil {
		return err
	}

	err = m.write(filePath)
	if err != nil {
		return err
	}

	return nil
}

// All returns all the entries from metadata
func (m *Metadata) All() ([]Entry, error) {
	err := m.read(filePath)
	if err != nil {
		return m.Entries, err
	}

	return m.Entries, nil
}
