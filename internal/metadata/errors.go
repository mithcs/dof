package metadata

import "errors"

var (
	ErrEntryNotFound = errors.New("entry not found")
	ErrPathNotFound  = errors.New("path not found")
)
