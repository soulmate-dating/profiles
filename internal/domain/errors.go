package domain

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrForbidden       = errors.New("forbidden")
	ErrIDAlreadyExists = errors.New("id already exists")
	ErrNotUnique       = errors.New("entity is not unique")
)
