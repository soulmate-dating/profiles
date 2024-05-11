package models

import (
	"github.com/google/uuid"
	"time"
)

type Profile struct {
	UserId           uuid.UUID `db:"user_id"`
	FirstName        string    `db:"first_name,omitempty" validate:"required"`
	LastName         string    `db:"last_name"`
	BirthDate        time.Time `db:"birth_date"`
	Sex              string    `db:"sex" validate:"custom=validateEnum,validValues=man,woman"`
	PreferredPartner string    `db:"preferred_partner" validate:"custom=validateEnum,validValues=man,woman,anyone"`
	Intention        string    `db:"intention,omitempty" validate:"custom=validateEnum,validValues=life partner,long-term relationship,short-term relationship,friendship,figuring it out,prefer not to say"`
	Height           uint32    `db:"height,omitempty"`
	HasChildren      bool      `db:"has_children,omitempty"`
	FamilyPlans      string    `db:"family_plans,omitempty" validate:"custom=validateEnum,validValues=don't want children,want children,open to children,not sure yet,prefer not to say"`
	Location         string    `db:"location,omitempty"`
	DrinksAlcohol    string    `db:"drinks_alcohol,omitempty" validate:"custom=validateEnum,validValues=no,sometimes,yes,prefer not to say"`
	Smokes           string    `db:"smokes,omitempty" validate:"custom=validateEnum,validValues=no,sometimes,yes,prefer not to say"`
}
