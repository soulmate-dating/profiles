package models

import "time"

type Profile struct {
	UserId           string    `db:"user_id"`
	FirstName        string    `db:"first_name,omitempty"`
	LastName         string    `db:"last_name"`
	BirthDate        time.Time `db:"birth_date"`
	Sex              string    `db:"sex,omitempty"`
	PreferredPartner string    `db:"preferred_partner,omitempty"`
	Intention        string    `db:"intention,omitempty"`
	Height           uint32    `db:"height,omitempty"`
	HasChildren      bool      `db:"has_children,omitempty"`
	FamilyPlans      string    `db:"family_plans,omitempty"`
	Location         string    `db:"location,omitempty"`
	DrinksAlcohol    string    `db:"drinks_alcohol,omitempty"`
	Smokes           string    `db:"smokes,omitempty"`
}
