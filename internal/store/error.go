package store

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	FailedPrecondition = errors.New("failed precondition")
)
