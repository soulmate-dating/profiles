package models

type Profile struct {
	UserId           string `json:"user_id"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name"`
	BirthDate        string `json:"birth_date"`
	Sex              string `json:"sex,omitempty"`
	PreferredPartner string `json:"preferred_partner,omitempty"`
	Intention        string `json:"intention,omitempty"`
	Height           uint32 `json:"height,omitempty"`
	HasChildren      bool   `json:"has_children,omitempty"`
	FamilyPlans      string `json:"family_plans,omitempty"`
	Location         string `json:"location,omitempty"`
	EducationLevel   string `json:"education_level,omitempty"`
	DrinksAlcohol    string `json:"drinks_alcohol,omitempty"`
	SmokesCigarettes string `json:"smokes_cigarettes,omitempty"`
}
