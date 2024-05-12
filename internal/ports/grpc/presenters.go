package grpc

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"

	"github.com/soulmate-dating/profiles/internal/domain"
)

func ProfileSuccessResponse(p *domain.Profile) *ProfileResponse {
	return &ProfileResponse{
		Id: p.UserId.String(),
		PersonalInfo: &PersonalInfo{
			FirstName:        p.FirstName,
			LastName:         p.LastName,
			BirthDate:        p.BirthDate.Format(domain.DateLayout),
			Sex:              p.Sex,
			PreferredPartner: p.PreferredPartner,
			Intention:        p.Intention,
			Height:           p.Height,
			HasChildren:      p.HasChildren,
			FamilyPlans:      p.FamilyPlans,
			Location:         p.Location,
			DrinksAlcohol:    p.DrinksAlcohol,
			Smokes:           p.Smokes,
			ProfilePicLink:   p.MainPicLink,
		},
	}
}

func GetMultipleProfilesSuccessResponse(profiles []domain.Profile) *MultipleProfilesResponse {
	res := make([]*ProfileResponse, len(profiles))
	for i, p := range profiles {
		res[i] = &ProfileResponse{
			Id: p.UserId.String(),
			PersonalInfo: &PersonalInfo{
				FirstName:        p.FirstName,
				LastName:         p.LastName,
				BirthDate:        p.BirthDate.Format(domain.DateLayout),
				Sex:              p.Sex,
				PreferredPartner: p.PreferredPartner,
				Intention:        p.Intention,
				Height:           p.Height,
				HasChildren:      p.HasChildren,
				FamilyPlans:      p.FamilyPlans,
				Location:         p.Location,
				DrinksAlcohol:    p.DrinksAlcohol,
				Smokes:           p.Smokes,
				ProfilePicLink:   p.MainPicLink,
			},
		}
	}
	return &MultipleProfilesResponse{Profiles: res}
}

func PromptsSuccessResponse(userId string, prompts []domain.Prompt) *PromptsResponse {
	res := make([]*Prompt, len(prompts))
	for i, p := range prompts {
		res[i] = &Prompt{
			Id:       p.ID.String(),
			Question: p.Question,
			Content:  p.Content,
			Position: p.Position,
			Type:     string(p.Type),
		}
	}
	return &PromptsResponse{UserId: userId, Prompts: res}
}

func SinglePromptSuccessResponse(p *domain.Prompt) *SinglePromptResponse {
	return &SinglePromptResponse{
		UserId: p.UserId.String(),
		Prompt: &Prompt{
			Id:       p.ID.String(),
			Question: p.Question,
			Content:  p.Content,
			Position: p.Position,
			Type:     string(p.Type),
		},
	}
}

func FullProfileSuccessResponse(fp *domain.FullProfile) *FullProfileResponse {
	prompts := fp.Prompts
	res := make([]*Prompt, len(prompts))
	for i, p := range prompts {
		res[i] = &Prompt{
			Id:       p.ID.String(),
			Question: p.Question,
			Content:  p.Content,
			Position: p.Position,
			Type:     string(p.Type),
		}
	}
	profile := fp.Profile
	return &FullProfileResponse{
		UserId: profile.UserId.String(),
		PersonalInfo: &PersonalInfo{
			FirstName:        profile.FirstName,
			LastName:         profile.LastName,
			BirthDate:        profile.BirthDate.Format(domain.DateLayout),
			Sex:              profile.Sex,
			PreferredPartner: profile.PreferredPartner,
			Intention:        profile.Intention,
			Height:           profile.Height,
			HasChildren:      profile.HasChildren,
			FamilyPlans:      profile.FamilyPlans,
			Location:         profile.Location,
			DrinksAlcohol:    profile.DrinksAlcohol,
			Smokes:           profile.Smokes,
			ProfilePicLink:   profile.MainPicLink,
		},
		Prompts: res,
	}
}

func mapCreateProfileRequest(request *CreateProfileRequest) (*domain.Profile, error) {
	info := request.GetPersonalInfo()
	userId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, err
	}
	birthDate, err := domain.ParseDate(info.GetBirthDate())
	if err != nil {
		return nil, err
	}
	return &domain.Profile{
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
	case errors.Is(err, domain.ErrNotFound):
		return codes.NotFound
	case errors.Is(err, domain.ErrNotUnique) || errors.Is(err, domain.ErrIDAlreadyExists):
		return codes.AlreadyExists
	case errors.Is(err, domain.ErrAddPromptsOnEmptyProfile):
		return codes.FailedPrecondition
	case errors.Is(err, domain.ErrForbidden):
		return codes.PermissionDenied
	}
	return codes.Internal
}
