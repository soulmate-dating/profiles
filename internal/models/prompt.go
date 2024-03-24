package models

type Prompt struct {
	UID      string `db:"uid"`
	UserId   string `db:"user_id"`
	Question string `db:"question"`
	Answer   string `db:"answer"`
	Position int32  `db:"position"`
}
