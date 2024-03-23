package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/soulmate-dating/profiles.git/internal/adapters/postgres"
	"github.com/soulmate-dating/profiles.git/internal/models"
)

var (
	ErrForbidden = fmt.Errorf("forbidden")
)

type ProfileApp interface {
	CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error)
	GetProfile(ctx context.Context, id string) (*models.Profile, error)
}

type App interface {
	ProfileApp
}

type ProfileRepository interface {
	CreateProfile(ctx context.Context, p *models.Profile) error
	GetProfileByID(ctx context.Context, id string) (models.Profile, error)
}

type Repository interface {
	ProfileRepository
}

type Application struct {
	repository Repository
}

func (a Application) CreateProfile(ctx context.Context, profile *models.Profile) (*models.Profile, error) {
	err := a.repository.CreateProfile(ctx, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (a Application) GetProfile(ctx context.Context, id string) (*models.Profile, error) {
	p, err := a.repository.GetProfileByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func NewApp(conn *pgxpool.Pool) App {
	repo := postgres.NewRepo(conn)
	return &Application{repository: repo}
}
