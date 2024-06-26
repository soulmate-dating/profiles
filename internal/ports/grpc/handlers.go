package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/soulmate-dating/profiles/internal/domain"
	"google.golang.org/grpc/codes"
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
	userId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	profile, err := s.app.GetProfile(ctx, userId)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return ProfileSuccessResponse(profile), nil
}

func (s *ProfileService) UpdateProfile(ctx context.Context, request *UpdateProfileRequest) (*ProfileResponse, error) {
	info := request.GetPersonalInfo()
	birthDate, err := domain.ParseDate(info.GetBirthDate())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	userId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	p := domain.Profile{
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
	profile, err := s.app.UpdateProfile(ctx, p)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return ProfileSuccessResponse(profile), nil
}

func (s *ProfileService) GetMultipleProfiles(ctx context.Context, request *GetMultipleProfilesRequest) (*MultipleProfilesResponse, error) {
	userIDs := make([]uuid.UUID, len(request.GetIds()))
	for i, id := range request.GetIds() {
		userId, err := uuid.Parse(id)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		userIDs[i] = userId
	}
	profiles, err := s.app.GetMultipleProfiles(ctx, userIDs)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return GetMultipleProfilesSuccessResponse(profiles), nil
}

func (s *ProfileService) GetRandomProfilePreferredByUser(ctx context.Context, request *GetRandomProfilePreferredByUserRequest) (*FullProfileResponse, error) {
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	profile, err := s.app.GetRandomProfilePreferredByUser(ctx, userId)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return FullProfileSuccessResponse(profile), nil
}

func (s *ProfileService) GetFullProfile(ctx context.Context, request *GetProfileRequest) (*FullProfileResponse, error) {
	userId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	profile, err := s.app.GetFullProfile(ctx, userId)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return FullProfileSuccessResponse(profile), nil
}

func (s *ProfileService) GetPrompts(ctx context.Context, request *GetPromptsRequest) (*PromptsResponse, error) {
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	prompts, err := s.app.GetPrompts(ctx, userId)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return PromptsSuccessResponse(request.GetUserId(), prompts), nil
}

func (s *ProfileService) AddPrompts(ctx context.Context, request *AddPromptsRequest) (*PromptsResponse, error) {
	prompts := make([]domain.Prompt, len(request.GetPrompts()))
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	for i, p := range request.GetPrompts() {
		prompts[i] = domain.Prompt{
			UserId:   userId,
			Question: p.GetQuestion(),
			Content:  p.GetContent(),
			Position: p.GetPosition(),
			Type:     domain.Text,
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
	p := domain.Prompt{
		ID:       promptId,
		UserId:   userId,
		Question: promptInfo.GetQuestion(),
		Content:  promptInfo.GetContent(),
		Position: promptInfo.GetPosition(),
		Type:     domain.ContentType(promptInfo.GetType()),
	}
	prompt, err := s.app.UpdatePrompt(ctx, p)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return SinglePromptSuccessResponse(prompt), nil
}

func (s *ProfileService) UpdatePromptsPositions(ctx context.Context, request *UpdatePromptsPositionsRequest) (*PromptsResponse, error) {
	prompts := make([]domain.Prompt, len(request.GetPromptPositions()))
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}

	for i, p := range request.GetPromptPositions() {
		promptId, err := uuid.Parse(p.GetId())
		if err != nil {
			return nil, status.Error(GetErrorCode(err), err.Error())
		}
		prompts[i] = domain.Prompt{
			ID:       promptId,
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

func (s *ProfileService) AddFilePrompt(ctx context.Context, request *AddFilePromptRequest) (*SinglePromptResponse, error) {
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	filePrompt := domain.FilePrompt{
		UserId:   userId,
		Question: request.GetQuestion(),
		Content:  request.GetContent(),
		Position: request.GetPosition(),
		Type:     domain.Image,
	}
	prompt, err := s.app.AddFilePrompt(ctx, filePrompt)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return SinglePromptSuccessResponse(prompt), nil
}

func (s *ProfileService) UpdateFilePrompt(ctx context.Context, request *UpdateFilePromptRequest) (*SinglePromptResponse, error) {
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	promptId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	filePrompt := domain.FilePrompt{
		ID:       promptId,
		UserId:   userId,
		Question: request.GetQuestion(),
		Content:  request.GetContent(),
		Position: request.GetPosition(),
		Type:     domain.Image,
	}
	prompt, err := s.app.UpdateFilePrompt(ctx, filePrompt)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return SinglePromptSuccessResponse(prompt), nil
}

func (s *ProfileService) DeletePrompt(ctx context.Context, request *DeletePromptRequest) (*SinglePromptResponse, error) {
	userId, err := uuid.Parse(request.GetUserId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	promptId, err := uuid.Parse(request.GetId())
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	prompt, err := s.app.DeletePrompt(ctx, userId, promptId)
	if err != nil {
		return nil, status.Error(GetErrorCode(err), err.Error())
	}
	return SinglePromptSuccessResponse(prompt), nil
}
