package models

import "github.com/google/uuid"

type ContentType string

const (
	Image ContentType = "image"
	Text  ContentType = "text"
)

type Prompt struct {
	ID       uuid.UUID   `db:"id"`
	UserId   uuid.UUID   `db:"user_id"`
	Question string      `db:"question"`
	Content  string      `db:"content"`
	Position int32       `db:"position" validate:"min=0,max=10"`
	Type     ContentType `db:"type" validate:"custom=validateEnum,validValues=image,text"`
}

type FilePrompt struct {
	ID       uuid.UUID   `db:"id"`
	UserId   uuid.UUID   `db:"user_id"`
	Question string      `db:"question"`
	Content  []byte      `db:"content"`
	Position int32       `db:"position"`
	Type     ContentType `db:"type"  validate:"custom=validateEnum,validValues=image,text"`
}
