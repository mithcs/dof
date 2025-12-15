package metadata

import (
	"encoding/json"
	"os"
)

// read reads the metadata file from given filepath
func (m *Metadata) read(filepath string) error {
	input, err := os.ReadFile(filepath)
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

// write writes to the metadata file to given filepath
func (m *Metadata) write(filepath string) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath, b, 0666)
	if err != nil {
		return err
	}

	return nil
}
