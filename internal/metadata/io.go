package metadata

import (
	"encoding/json"
	"os"
)

// read reads the metadata file
func (m *Metadata) read(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// exit if file is empty
	if len(input) == 0 {
		return nil
	}

	err = json.Unmarshal(input, m)
	if err != nil {
		return err
	}

	return nil
}

// write writes to the metadata file
func (m *Metadata) write(filename string) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0666)
	if err != nil {
		return err
	}

	return nil
}
