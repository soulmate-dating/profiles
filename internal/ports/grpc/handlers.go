package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/soulmate-dating/profiles/internal/models"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *ProfileService) CreateProfile(ctx context.Context, request *CreateProfileRequest) (*ProfileResponse, error) {
	profile, err := mapCreateProfileRequest(request)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	profile, err = s.app.CreateProfile(ctx, profile)
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
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	userId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	p := models.Profile{
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
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	for i, p := range request.GetPrompts() {
		prompts[i] = models.Prompt{
			UserId:   userId,
			Question: p.GetQuestion(),
			Answer:   p.GetAnswer(),
			Position: p.GetPosition(),
		}
	}

	prompts, err = s.app.AddPrompts(ctx, prompts)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return PromptsSuccessResponse(request.GetUserId(), prompts), nil
}

func (s *ProfileService) UpdatePrompt(ctx context.Context, request *UpdatePromptRequest) (*SinglePromptResponse, error) {
	promptInfo := request.GetPrompt()
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	promptId, err := uuid.Parse(promptInfo.GetId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	p := models.Prompt{
		ID:       promptId,
		UserId:   userId,
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
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	for i, p := range request.GetPromptPositions() {
		prompts[i] = models.Prompt{
			UserId:   userId,
			Position: p.GetPosition(),
		}
	}
	prompts, err = s.app.UpdatePromptsPositions(ctx, prompts)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return PromptsSuccessResponse(request.GetUserId(), prompts), nil
}
