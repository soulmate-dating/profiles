package app

import (
	"context"
	"fmt"
	"github.com/soulmate-dating/profiles/internal/app/clients/media"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/soulmate-dating/profiles/internal/adapters/postgres"
	"github.com/soulmate-dating/profiles/internal/models"
)

var (
	ErrForbidden = fmt.Errorf("forbidden")
)

type App interface {
	CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	GetProfile(ctx context.Context, userId string) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	GetRandomProfilePreferredByUser(ctx context.Context, userId string) (*models.FullProfile, error)
	GetFullProfile(ctx context.Context, userId string) (*models.FullProfile, error)

	GetPrompts(ctx context.Context, userId string) ([]models.Prompt, error)
	AddPrompts(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error)
	UpdatePrompt(ctx context.Context, prompt *models.Prompt) (*models.Prompt, error)
	UpdatePromptsPositions(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error)
	GetMultipleProfiles(ctx context.Context, ids []string) ([]models.Profile, error)
	AddFilePrompt(ctx context.Context, prompt models.FilePrompt) (*models.Prompt, error)
}

type Repository interface {
	CreateProfile(ctx context.Context, p *models.Profile) error
	GetProfileByID(ctx context.Context, id string) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	GetMultipleProfilesByIDs(ctx context.Context, ids []string) ([]models.Profile, error)
	GetRandomProfileBySexAndPreference(
		ctx context.Context, requesterId uuid.UUID, preference models.Preference, sex string,
	) (*models.Profile, error)

	GetPromptsByUser(ctx context.Context, userId string) ([]models.Prompt, error)
	CreatePrompt(ctx context.Context, prompt models.Prompt) error
	UpdatePromptContent(ctx context.Context, prompt *models.Prompt) (*models.Prompt, error)
	UpdatePromptsPositions(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error)
}

type Application struct {
	repository  Repository
	mediaClient media.MediaServiceClient
}

func (a *Application) AddFilePrompt(ctx context.Context, filePrompt models.FilePrompt) (*models.Prompt, error) {
	response, err := a.mediaClient.UploadFile(ctx, &media.UploadFileRequest{
		ContentType: "image/png",
		Data:        filePrompt.Content,
	})
	if err != nil {
		return nil, err
	}

	prompt := models.Prompt{
		ID:       models.NewUID(),
		UserId:   filePrompt.UserId,
		Question: filePrompt.Question,
		Content:  response.GetLink(),
		Position: filePrompt.Position,
		Type:     filePrompt.Type,
	}
	err = a.repository.CreatePrompt(ctx, prompt)
	return &prompt, err
}

func (a *Application) GetFullProfile(ctx context.Context, userId string) (*models.FullProfile, error) {
	profile, err := a.repository.GetProfileByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	prompts, err := a.repository.GetPromptsByUser(ctx, profile.UserId.String())
	if err != nil {
		return nil, err
	}

	return &models.FullProfile{
		Profile: *profile,
		Prompts: prompts,
	}, nil
}

func (a *Application) GetRandomProfilePreferredByUser(ctx context.Context, userId string) (*models.FullProfile, error) {
	profile, err := a.repository.GetProfileByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	recommendedProfile, err := a.repository.GetRandomProfileBySexAndPreference(
		ctx, profile.UserId, models.Preference(profile.PreferredPartner), profile.Sex,
	)
	if err != nil {
		return nil, err
	}

	prompts, err := a.repository.GetPromptsByUser(ctx, recommendedProfile.UserId.String())
	if err != nil {
		return nil, err
	}

	return &models.FullProfile{
		Profile: *recommendedProfile,
		Prompts: prompts,
	}, nil
}

func (a *Application) GetMultipleProfiles(ctx context.Context, ids []string) ([]models.Profile, error) {
	p, err := a.repository.GetMultipleProfilesByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *Application) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	err := a.repository.CreateProfile(ctx, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (a *Application) GetProfile(ctx context.Context, userId string) (*models.Profile, error) {
	p, err := a.repository.GetProfileByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *Application) UpdateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	p, err := a.repository.UpdateProfile(ctx, profile)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *Application) GetPrompts(ctx context.Context, userId string) ([]models.Prompt, error) {
	prompts, err := a.repository.GetPromptsByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return prompts, nil
}

func (a *Application) AddPrompts(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error) {
	for i := range prompts {
		prompts[i].ID = models.NewUID()
		err := a.repository.CreatePrompt(ctx, prompts[i])
		if err != nil {
			return nil, err
		}
	}

	return prompts, nil
}

func (a *Application) UpdatePrompt(ctx context.Context, prompt *models.Prompt) (*models.Prompt, error) {
	prompt, err := a.repository.UpdatePromptContent(ctx, prompt)
	if err != nil {
		return nil, err
	}

	return prompt, nil
}

func (a *Application) UpdatePromptsPositions(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error) {
	prompts, err := a.repository.UpdatePromptsPositions(ctx, prompts)
	if err != nil {
		return nil, err
	}

	return prompts, nil
}

func NewApp(conn *pgxpool.Pool) App {
	repo := postgres.NewRepo(conn)
	mediaClient, err := media.NewServiceClient()
	if err != nil {
		log.Fatalf("could not connect to media service: %s", err.Error())
	}
	return &Application{repository: repo, mediaClient: mediaClient}
}
