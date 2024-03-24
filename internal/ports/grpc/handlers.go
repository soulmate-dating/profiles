package grpc

import (
	"context"
	"github.com/soulmate-dating/profiles.git/internal/models"
	"google.golang.org/grpc/status"
)

func (s *ProfileService) CreateProfile(ctx context.Context, request *CreateProfileRequest) (*ProfileResponse, error) {
	info := request.GetPersonalInfo()
	birthDate, err := models.ParseDate(info.GetBirthDate())
	p := models.Profile{
		UserId:           request.GetId(),
		FirstName:        info.GetFirstName(),
		LastName:         info.GetLastName(),
		BirthDate:        birthDate,
		Sex:              info.GetSex(),
		PreferredPartner: info.GetPreferredPartner(),
		Intention:        info.GetIntention(),
		Height:           info.GetHeight(),
		HasChildren:      info.GetHasChildren(),
		FamilyPlans:      info.GetFamilyPlans(),
		Location:         info.GetLocation(),
		DrinksAlcohol:    info.GetDrinksAlcohol(),
		Smokes:           info.GetSmokes(),
	}
	profile, err := s.app.CreateProfile(ctx, &p)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return ProfileSuccessResponse(profile), nil
}

func (s *ProfileService) GetProfile(ctx context.Context, request *GetProfileRequest) (*ProfileResponse, error) {
	profile, err := s.app.GetProfile(ctx, request.GetId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return ProfileSuccessResponse(profile), nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, request *UpdateProfileRequest) (*ProfileResponse, error) {
	info := request.GetPersonalInfo()
	birthDate, err := models.ParseDate(info.GetBirthDate())
	p := models.Profile{
		UserId:           request.GetId(),
		FirstName:        info.GetFirstName(),
		LastName:         info.GetLastName(),
		BirthDate:        birthDate,
		Sex:              info.GetSex(),
		PreferredPartner: info.GetPreferredPartner(),
		Intention:        info.GetIntention(),
		Height:           info.GetHeight(),
		HasChildren:      info.GetHasChildren(),
		FamilyPlans:      info.GetFamilyPlans(),
		Location:         info.GetLocation(),
		DrinksAlcohol:    info.GetDrinksAlcohol(),
		Smokes:           info.GetSmokes(),
	}
	profile, err := s.app.UpdateProfile(ctx, &p)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return ProfileSuccessResponse(profile), nil
}

func (s *ProfileService) GetPrompts(ctx context.Context, request *GetPromptsRequest) (*PromptsResponse, error) {
	prompts, err := s.app.GetPrompts(ctx, request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return PromptsSuccessResponse(request.GetUserId(), prompts), nil
}

func (s *ProfileService) AddPrompts(ctx context.Context, request *AddPromptsRequest) (*PromptsResponse, error) {
	prompts := make([]models.Prompt, len(request.GetPrompts()))
	for i, p := range request.GetPrompts() {
		prompts[i] = models.Prompt{
			UserId:   request.GetUserId(),
			Question: p.GetQuestion(),
			Answer:   p.GetAnswer(),
			Position: p.GetPosition(),
		}
	}

	prompts, err := s.app.AddPrompts(ctx, prompts)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return PromptsSuccessResponse(request.GetUserId(), prompts), nil
}

func (s *ProfileService) UpdatePrompt(ctx context.Context, request *UpdatePromptRequest) (*SinglePromptResponse, error) {
	promptInfo := request.GetPrompt()
	p := models.Prompt{
		UID:      promptInfo.GetId(),
		UserId:   request.GetUserId(),
		Question: promptInfo.GetQuestion(),
		Answer:   promptInfo.GetAnswer(),
		Position: promptInfo.GetPosition(),
	}
	prompt, err := s.app.UpdatePrompt(ctx, &p)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return SinglePromptSuccessResponse(prompt), nil
}

func (s *ProfileService) UpdatePromptsPositions(ctx context.Context, request *UpdatePromptsPositionsRequest) (*PromptsResponse, error) {
	prompts := make([]models.Prompt, len(request.GetPromptPositions()))
	for i, p := range request.GetPromptPositions() {
		prompts[i] = models.Prompt{
			UserId:   request.GetUserId(),
			Position: p.GetPosition(),
		}
	}
	prompts, err := s.app.UpdatePromptsPositions(ctx, prompts)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return PromptsSuccessResponse(request.GetUserId(), prompts), nil
}
