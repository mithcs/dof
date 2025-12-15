package metadata

// add appends given entry to metadata
func (m *Metadata) add(e Entry) error {
	m.Entries = append(m.Entries, e)
	return nil
}

// remove removes given entry from metadata
func (m *Metadata) remove(name string) error {
	for i, e := range m.Entries {
		if e.Name == name {
			m.Entries = append(m.Entries[:i], m.Entries[i+1:]...)
			return nil
		}
	}

	return ErrEntryNotFound
}

// get returns entry with specified name from metadata
func (m *Metadata) get(name string) (Entry, error) {
	for _, e := range m.Entries {
		if e.Name == name {
			return e, nil
		}
	}

	return Entry{}, ErrEntryNotFound
}

// update replaces existing entry with updated entry
func (m *Metadata) update(updated Entry) error {
	for i := range m.Entries {
		if m.Entries[i].Name == updated.Name {
			m.Entries[i] = updated
			return nil
		}
	}

	return ErrEntryNotFound
}
