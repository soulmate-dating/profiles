package grpc

import (
	"errors"
	"github.com/TobbyMax/validator"
	"github.com/soulmate-dating/profiles.git/internal/app"
	"github.com/soulmate-dating/profiles.git/internal/models"
	"google.golang.org/grpc/codes"
)

var ErrMissingArgument = errors.New("required argument is missing")

func ProfileSuccessResponse(p *models.Profile) *ProfileResponse {
	return &ProfileResponse{
		Id: p.UserId,
		PersonalInfo: &PersonalInfo{
			FirstName:        p.FirstName,
			LastName:         p.LastName,
			BirthDate:        p.BirthDate,
			Sex:              p.Sex,
			PreferredPartner: p.PreferredPartner,
			Intention:        p.Intention,
			Height:           p.Height,
			HasChildren:      p.HasChildren,
			FamilyPlans:      p.FamilyPlans,
			Location:         p.Location,
			EducationLevel:   p.EducationLevel,
			DrinksAlcohol:    p.DrinksAlcohol,
			SmokesCigarettes: p.SmokesCigarettes,
		},
	}
}

func GetErrorCode(err error) codes.Code {
	switch {
	case errors.As(err, &validator.ValidationErrors{}):
		return codes.InvalidArgument
	case errors.Is(err, app.ErrForbidden):
		return codes.PermissionDenied
	}
	return codes.Internal
}
