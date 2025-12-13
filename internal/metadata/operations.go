package metadata

// add adds entry to metadata
func (m *Metadata) add(e Entry) error {
	m.Entries = append(m.Entries, e)
	return nil
}

// remove removes entry from metadata
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

// update replaces existing entry with new entry (using name)
func (m *Metadata) update(new Entry) error {
	for i := range m.Entries {
		if m.Entries[i].Name == new.Name {
			m.Entries[i] = new
			return nil
		}
	}

	return ErrEntryNotFound
}
