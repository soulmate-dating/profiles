package grpc

import (
	"errors"
	"github.com/TobbyMax/validator"
	"github.com/google/uuid"
	"github.com/soulmate-dating/profiles/internal/app"
	"github.com/soulmate-dating/profiles/internal/models"
	"google.golang.org/grpc/codes"
	"strings"
)

var ErrMissingArgument = errors.New("required argument is missing")

func ProfileSuccessResponse(p *models.Profile) *ProfileResponse {
	return &ProfileResponse{
		Id: p.UserId.String(),
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

func GetMultipleProfilesSuccessResponse(profiles []models.Profile) *MultipleProfilesResponse {
	res := make([]*ProfileResponse, len(profiles))
	for i, p := range profiles {
		res[i] = &ProfileResponse{
			Id: p.UserId.String(),
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
	return &MultipleProfilesResponse{Profiles: res}
}

func PromptsSuccessResponse(userId string, prompts []models.Prompt) *PromptsResponse {
	res := make([]*Prompt, len(prompts))
	for i, p := range prompts {
		res[i] = &Prompt{
			Id:       p.ID.String(),
			Question: p.Question,
			Answer:   p.Answer,
			Position: p.Position,
		}
	}
	return &PromptsResponse{UserId: userId, Prompts: res}
}

func SinglePromptSuccessResponse(p *models.Prompt) *SinglePromptResponse {
	return &SinglePromptResponse{
		UserId: p.UserId.String(),
		Prompt: &Prompt{
			Id:       p.ID.String(),
			Question: p.Question,
			Answer:   p.Answer,
			Position: p.Position,
		},
	}
}

func FullProfileSuccessResponse(fp *models.FullProfile) *FullProfileResponse {
	prompts := fp.Prompts
	res := make([]*Prompt, len(prompts))
	for i, p := range prompts {
		res[i] = &Prompt{
			Id:       p.ID.String(),
			Question: p.Question,
			Answer:   p.Answer,
			Position: p.Position,
		}
	}
	profile := fp.Profile
	return &FullProfileResponse{
		UserId: profile.UserId.String(),
		PersonalInfo: &PersonalInfo{
			FirstName:        profile.FirstName,
			LastName:         profile.LastName,
			BirthDate:        profile.BirthDate.Format(models.DateLayout),
			Sex:              profile.Sex,
			PreferredPartner: profile.PreferredPartner,
			Intention:        profile.Intention,
			Height:           profile.Height,
			HasChildren:      profile.HasChildren,
			FamilyPlans:      profile.FamilyPlans,
			Location:         profile.Location,
			DrinksAlcohol:    profile.DrinksAlcohol,
			Smokes:           profile.Smokes,
		},
		Prompts: res,
	}
}

func mapCreateProfileRequest(request *CreateProfileRequest) (*models.Profile, error) {
	info := request.GetPersonalInfo()
	userId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	birthDate, err := models.ParseDate(info.GetBirthDate())
	if err != nil {
		return nil, err
	}
	return &models.Profile{
		UserId:           userId,
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
