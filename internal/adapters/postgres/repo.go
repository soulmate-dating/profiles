package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/soulmate-dating/profiles/internal/models"
)

type Repo struct {
	pool        *pgxpool.Pool
	mapProfiles func(row pgx.CollectableRow) (models.Profile, error)
	mapPrompts  func(row pgx.CollectableRow) (models.Prompt, error)
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{
		pool:        pool,
		mapProfiles: pgx.RowToStructByName[models.Profile],
		mapPrompts:  pgx.RowToStructByName[models.Prompt],
	}
}

func (r *Repo) CreateProfile(ctx context.Context, p *models.Profile) error {
	var args []any
	args = append(args,
		p.UserId, p.FirstName, p.LastName, p.BirthDate,
		p.Sex, p.PreferredPartner, p.Intention, p.Height,
		p.HasChildren, p.FamilyPlans, p.Location,
		p.DrinksAlcohol, p.Smokes,
	)
	if _, err := r.pool.Exec(ctx, createProfileQuery, args...); err != nil {
		return fmt.Errorf("create profile: %w", err)
	}
	return nil
}

func (r *Repo) GetProfileByID(ctx context.Context, id string) (*models.Profile, error) {
	rows, err := r.pool.Query(ctx, getProfileByIDQuery, id)
	if err != nil {
		return nil, fmt.Errorf("get profile by id: %w", err)
	}
	profile, err := pgx.CollectOneRow(rows, r.mapProfiles)
	if err != nil {
		return nil, fmt.Errorf("map profile: %w", err)
	}
	return &profile, nil
}

func (r *Repo) UpdateProfile(ctx context.Context, p *models.Profile) (*models.Profile, error) {
	var args []any
	args = append(args,
		p.UserId, p.FirstName, p.LastName, p.BirthDate,
		p.Sex, p.PreferredPartner, p.Intention, p.Height,
		p.HasChildren, p.FamilyPlans, p.Location,
		p.DrinksAlcohol, p.Smokes,
	)
	rows, err := r.pool.Query(ctx, updateProfileQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("update profile: %w", err)
	}
	profile, err := pgx.CollectOneRow(rows, r.mapProfiles)
	if err != nil {
		return nil, fmt.Errorf("map profile: %w", err)
	}
	return &profile, nil
}

func (r *Repo) GetPromptsByUser(ctx context.Context, userId string) ([]models.Prompt, error) {
	rows, err := r.pool.Query(ctx, getPromptsQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("get prmpts by id: %w", err)
	}
	prompts, err := pgx.CollectRows(rows, r.mapPrompts)
	if err != nil {
		return nil, fmt.Errorf("map prompts: %w", err)
	}
	return prompts, nil
}

func (r *Repo) CreatePrompt(ctx context.Context, prompt models.Prompt) error {
	var args []any
	args = append(args,
		prompt.UID, prompt.UserId, prompt.Question, prompt.Answer, prompt.Position,
	)
	if _, err := r.pool.Exec(ctx, createPromptQuery, args...); err != nil {
		return fmt.Errorf("create prompt: %w", err)
	}
	return nil
}

func (r *Repo) UpdatePromptContent(ctx context.Context, prompt *models.Prompt) (*models.Prompt, error) {
	var args []any
	args = append(args,
		prompt.UID, prompt.Question, prompt.Answer,
	)
	rows, err := r.pool.Query(ctx, updatePromptQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("update prompt: %w", err)
	}
	p, err := pgx.CollectOneRow(rows, r.mapPrompts)
	if err != nil {
		return nil, fmt.Errorf("map prompt: %w", err)
	}
	return &p, nil
}

func (r *Repo) UpdatePromptsPositions(ctx context.Context, prompts []models.Prompt) ([]models.Prompt, error) {
	//TODO implement me
	panic("implement me")
}
