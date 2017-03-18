package core

import (
	"errors"
	"fmt"
)

var (
	// ErrTypeAssertion is the error for a type assertion failure.
	ErrTypeAssertion = errors.New("Type assertion failed")
)

// EntityNotFound is the error used to capture the case where the specified there are no entities of the specified kind and ID.
type EntityNotFound struct {
	id   string
	kind string
}

// NewEntityNotFound returns a new instance of EntityNotFound.
func NewEntityNotFound(id, kind string) error {
	return &EntityNotFound{id: id, kind: kind}
}

// Error returns the string representation of e.
func (e *EntityNotFound) Error() string {
	return fmt.Sprintf("Unable to find %s with ID %q", e.kind, e.id)
}
