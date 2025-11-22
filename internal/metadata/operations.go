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
