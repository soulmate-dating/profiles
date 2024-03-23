package grpc

import (
	"context"
	"github.com/soulmate-dating/profiles.git/internal/models"
	"google.golang.org/grpc/status"
)

func (s *ProfileService) CreateProfile(ctx context.Context, request *CreateProfileRequest) (*ProfileResponse, error) {
	info := request.GetPersonalInfo()
	p := models.Profile{
		UserId:           request.GetId(),
		FirstName:        info.FirstName,
		LastName:         info.LastName,
		BirthDate:        info.BirthDate,
		Sex:              info.Sex,
		PreferredPartner: info.PreferredPartner,
		Intention:        info.Intention,
		Height:           info.Height,
		HasChildren:      info.HasChildren,
		FamilyPlans:      info.FamilyPlans,
		Location:         info.Location,
		EducationLevel:   info.EducationLevel,
		DrinksAlcohol:    info.DrinksAlcohol,
		SmokesCigarettes: info.SmokesCigarettes,
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
	//TODO implement me
	panic("implement me")
}

func (s *ProfileService) GetPrompts(ctx context.Context, request *GetPromptsRequest) (*PromptsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProfileService) AddPrompts(ctx context.Context, request *AddPromptsRequest) (*PromptsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProfileService) UpdatePrompt(ctx context.Context, request *UpdatePromptRequest) (*UpdatePromptResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProfileService) ReorderPrompts(ctx context.Context, request *ReorderPromptsRequest) (*ReorderPromptsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProfileService) mustEmbedUnimplementedProfileServiceServer() {
	//TODO implement me
	panic("implement me")
}
