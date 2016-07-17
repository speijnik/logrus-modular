package modular

import "errors"

var (
	// ErrChildExists denotes that a child logger already exists
	ErrChildExists = errors.New("Child logger exists")
	// ErrChildNotFound denotes that a child logger was not found
	ErrChildNotFound = errors.New("Child logger not found")
)
