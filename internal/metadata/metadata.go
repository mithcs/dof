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

// Get returns entry with specified name from metadata
func (m *Metadata) Get(name string) (Entry, error) {
	err := read(m, files.MetadataPath(filename))
	if err != nil {
		return Entry{}, err
	}

	return get(m, name)
}

// Update replaces existing entry with new entry (using Name) in Metadata
func (m *Metadata) Update(new Entry) error {
	err := read(m, files.MetadataPath(filename))
	if err != nil {
		return err
	}

	err = update(m, new)
	if err != nil {
		return err
	}

	err = write(m, files.MetadataPath(filename))
	if err != nil {
		return err
	}

	return nil
}
