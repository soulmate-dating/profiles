package grpc

import (
	"errors"
	"github.com/TobbyMax/validator"
	"github.com/soulmate-dating/profiles/internal/app"
	"github.com/soulmate-dating/profiles/internal/models"
	"google.golang.org/grpc/codes"
	"strings"
)

var ErrMissingArgument = errors.New("required argument is missing")

func ProfileSuccessResponse(p *models.Profile) *ProfileResponse {
	return &ProfileResponse{
		Id: p.UserId,
		PersonalInfo: &PersonalInfo{
			FirstName:        p.FirstName,
			LastName:         p.LastName,
			BirthDate:        p.BirthDate.Format(models.DateLayout),
			Sex:              p.Sex,
			PreferredPartner: p.PreferredPartner,
			Intention:        p.Intention,
			Height:           p.Height,
			HasChildren:      p.HasChildren,
			FamilyPlans:      p.FamilyPlans,
			Location:         p.Location,
			DrinksAlcohol:    p.DrinksAlcohol,
			Smokes:           p.Smokes,
		},
	}
}

func PromptsSuccessResponse(userId string, prompts []models.Prompt) *PromptsResponse {
	var res []*Prompt
	for _, p := range prompts {
		res = append(res, &Prompt{
			Id:       p.UID,
			Question: p.Question,
			Answer:   p.Answer,
			Position: p.Position,
		})
	}
	return &PromptsResponse{UserId: userId, Prompts: res}
}

func SinglePromptSuccessResponse(p *models.Prompt) *SinglePromptResponse {
	return &SinglePromptResponse{
		UserId: p.UserId,
		Prompt: &Prompt{
			Id:       p.UID,
			Question: p.Question,
			Answer:   p.Answer,
			Position: p.Position,
		},
	}
}

func mapCreateProfileRequest(request *CreateProfileRequest) (*models.Profile, error) {
	info := request.GetPersonalInfo()
	birthDate, err := models.ParseDate(info.GetBirthDate())
	if err != nil {
		return nil, err
	}
	return &models.Profile{
		UserId:           request.GetId(),
		FirstName:        info.GetFirstName(),
		LastName:         info.GetLastName(),
		BirthDate:        birthDate,
		Sex:              strings.ToLower(info.GetSex()),
		PreferredPartner: strings.ToLower(info.GetPreferredPartner()),
		Intention:        strings.ToLower(info.GetIntention()),
		Height:           info.GetHeight(),
		HasChildren:      info.GetHasChildren(),
		FamilyPlans:      strings.ToLower(info.GetFamilyPlans()),
		Location:         info.GetLocation(),
		DrinksAlcohol:    strings.ToLower(info.GetDrinksAlcohol()),
		Smokes:           strings.ToLower(info.GetSmokes()),
	}, nil
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
