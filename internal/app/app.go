package app

import (
	"context"
	"fmt"
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

	GetPrompts(ctx context.Context, userId string) ([]models.Prompt, error)
	AddPrompts(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error)
	UpdatePrompt(ctx context.Context, prompt *models.Prompt) (*models.Prompt, error)
	UpdatePromptsPositions(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error)
	GetMultipleProfiles(ctx context.Context, ids []string) ([]models.Profile, error)
}

type Repository interface {
	CreateProfile(ctx context.Context, p *models.Profile) error
	GetProfileByID(ctx context.Context, id string) (*models.Profile, error)
	UpdateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)

	GetPromptsByUser(ctx context.Context, userId string) ([]models.Prompt, error)
	CreatePrompt(ctx context.Context, prompt models.Prompt) error
	UpdatePromptContent(ctx context.Context, prompt *models.Prompt) (*models.Prompt, error)
	UpdatePromptsPositions(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error)
	GetMultipleProfilesByIDs(ctx context.Context, ids []string) ([]models.Profile, error)
}

type Application struct {
	repository Repository
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
	return &Application{repository: repo}
}
