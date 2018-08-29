package common

import "github.com/pkg/errors"

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrDoesNotExist  = errors.New("object does not exist")
	ErrCanNotDelete  = errors.New("object should not delete ")
)
