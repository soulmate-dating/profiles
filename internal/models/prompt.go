package models

import "github.com/google/uuid"

type Prompt struct {
	ID       uuid.UUID `db:"id"`
	UserId   uuid.UUID `db:"user_id"`
	Question string    `db:"question"`
	Answer   string    `db:"answer"`
	Position int32     `db:"position"`
}
