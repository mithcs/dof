package metadata

// add adds entry to metadata
func add(m *Metadata, e Entry) error {
	m.Entries = append(m.Entries, e)
	return nil
}

// remove removes entry from metadata
func remove(m *Metadata, name string) error {
	for i, e := range m.Entries {
		if e.Name == name {
			m.Entries = append(m.Entries[:i], m.Entries[i+1:]...)
			return nil
		}
	}

	return ErrEntryNotFound
}

// get returns entry with specified name from metadata
func get(m *Metadata, name string) (Entry, error) {
	for _, e := range m.Entries {
		if e.Name == name {
			return e, nil
		}
	}

	return Entry{}, ErrEntryNotFound
}

// update replaces existing entry with new entry (using name)
func update(m *Metadata, new Entry) error {
	for i := range m.Entries {
		if m.Entries[i].Name == new.Name {
			m.Entries[i] = new
			return nil
		}
	}

	return ErrEntryNotFound
}
