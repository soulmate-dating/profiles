package domain

import "github.com/google/uuid"

func NewUID() uuid.UUID {
	return uuid.New()
}
