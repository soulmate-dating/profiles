package domain

import (
	"errors"
)

var (
	ErrNotFound                 = errors.New("not found")
	ErrForbidden                = errors.New("forbidden")
	ErrIDAlreadyExists          = errors.New("id already exists")
	ErrNotUnique                = errors.New("entity is not unique")
	ErrAddPromptsOnEmptyProfile = errors.New("create profile before adding prompts")
	ErrCannotDeleteProfilePic   = errors.New("cannot delete profile picture")
)
