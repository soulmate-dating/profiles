package models

import "github.com/google/uuid"

func NewUID() string {
	return uuid.New().String()
}
