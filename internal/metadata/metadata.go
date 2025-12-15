package metadata

import (
	"github.com/mithcs/dof/internal/files"
)

type Method string

type Entry struct {
	Name   string   `json:"name"`
	Paths  []string `json:"paths"`
	Method Method   `json:"method"`
}

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

// Add adds entry to metadata
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

// Remove removes entry from metadata
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

// Get returns entry with specified name from metadata
func (m *Metadata) Get(name string) (Entry, error) {
	err := m.read(filePath)
	if err != nil {
		return Entry{}, err
	}

	return m.get(name)
}

// Update replaces existing entry with new entry (using Name) in Metadata
func (m *Metadata) Update(new Entry) error {
	err := m.read(filePath)
	if err != nil {
		return err
	}

	err = m.update(new)
	if err != nil {
		return err
	}

	err = m.write(filePath)
	if err != nil {
		return err
	}

	return nil
}

// All returns all the entries
func (m *Metadata) All() ([]Entry, error) {
	err := m.read(filePath)
	if err != nil {
		return m.Entries, err
	}

	return m.Entries, nil
}
