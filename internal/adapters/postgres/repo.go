package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/soulmate-dating/profiles/internal/domain"
)

type Repo struct {
	pool        ConnPool
	mapProfiles func(row pgx.CollectableRow) (domain.Profile, error)
	mapPrompts  func(row pgx.CollectableRow) (domain.Prompt, error)
}

func NewRepo(pool ConnPool) *Repo {
	return &Repo{
		pool:        pool,
		mapProfiles: pgx.RowToStructByName[domain.Profile],
		mapPrompts:  pgx.RowToStructByName[domain.Prompt],
	}
}

func (r *Repo) CreateProfile(ctx context.Context, p *domain.Profile) error {
	var args []any
	args = append(args,
		p.UserId, p.FirstName, p.LastName, p.BirthDate,
		p.Sex, p.PreferredPartner, p.Intention, p.Height,
		p.HasChildren, p.FamilyPlans, p.Location,
		p.DrinksAlcohol, p.Smokes,
	)
	if _, err := r.pool.GetTx(ctx).Exec(ctx, createProfileQuery, args...); err != nil {
		return fmt.Errorf("create profile: %w", err)
	}
	return nil
}

func (r *Repo) GetProfileByID(ctx context.Context, id uuid.UUID) (*domain.Profile, error) {
	rows, err := r.pool.GetTx(ctx).Query(ctx, getProfileByIDQuery, id)
	if err != nil {
		return nil, fmt.Errorf("get profile by id: %w", err)
	}
	profile, err := pgx.CollectOneRow(rows, r.mapProfiles)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("map profile: %w", err)
	}
	return &profile, nil
}

func (r *Repo) GetMultipleProfilesByIDs(ctx context.Context, userIds []uuid.UUID) ([]domain.Profile, error) {
	rows, err := r.pool.GetTx(ctx).Query(ctx, getMultipleProfilesByIDsQuery, userIds)
	if err != nil {
		return nil, fmt.Errorf("get profiles by id: %w", err)
	}
	prompts, err := pgx.CollectRows(rows, r.mapProfiles)
	if err != nil {
		return nil, fmt.Errorf("map profiles: %w", err)
	}
	return prompts, nil
}

func (r *Repo) GetRandomProfileBySexAndPreference(
	ctx context.Context, requesterId uuid.UUID, preference domain.Preference, sex string,
) (*domain.Profile, error) {
	pref1, pref2 := preference.Preferences()
	rows, err := r.pool.GetTx(ctx).Query(ctx, getRandomProfileBySexAndPreferenceQuery, requesterId, pref1, pref2, sex)
	if err != nil {
		return nil, fmt.Errorf("get profile by id: %w", err)
	}
	profile, err := pgx.CollectOneRow(rows, r.mapProfiles)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("map profile: %w", err)
	}
	return &profile, nil
}

func (r *Repo) UpdateProfile(ctx context.Context, p domain.Profile) (*domain.Profile, error) {
	var args []any
	args = append(args,
		p.UserId, p.FirstName, p.LastName, p.BirthDate,
		p.Sex, p.PreferredPartner, p.Intention, p.Height,
		p.HasChildren, p.FamilyPlans, p.Location,
		p.DrinksAlcohol, p.Smokes,
	)
	rows, err := r.pool.GetTx(ctx).Query(ctx, updateProfileQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("update profile: %w", err)
	}
	profile, err := pgx.CollectOneRow(rows, r.mapProfiles)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("map profile: %w", err)
	}
	return &profile, nil
}

func (r *Repo) GetPromptsByUser(ctx context.Context, userId uuid.UUID) ([]domain.Prompt, error) {
	rows, err := r.pool.GetTx(ctx).Query(ctx, getPromptsByUserQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("get prompts by id: %w", err)
	}
	prompts, err := pgx.CollectRows(rows, r.mapPrompts)
	if err != nil {
		return nil, fmt.Errorf("map prompts: %w", err)
	}
	return prompts, nil
}

func (r *Repo) CreatePrompt(ctx context.Context, prompt domain.Prompt) error {
	var args []any
	args = append(args,
		prompt.ID, prompt.UserId, prompt.Question, prompt.Content, prompt.Type, prompt.Position,
	)
	if _, err := r.pool.GetTx(ctx).Exec(ctx, createPromptQuery, args...); err != nil {
		return fmt.Errorf("create prompt: %w", err)
	}
	return nil
}

func (r *Repo) UpdatePromptContent(ctx context.Context, prompt domain.Prompt) (*domain.Prompt, error) {
	var args []any
	args = append(args,
		prompt.ID, prompt.Question, prompt.Content, prompt.Position,
	)
	rows, err := r.pool.GetTx(ctx).Query(ctx, updatePromptQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("update prompt: %w", err)
	}
	p, err := pgx.CollectOneRow(rows, r.mapPrompts)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("map prompt: %w", err)
	}
	return &p, nil
}

func (r *Repo) GetPromptByID(ctx context.Context, id uuid.UUID) (*domain.Prompt, error) {
	rows, err := r.pool.GetTx(ctx).Query(ctx, getPromptByIDQuery, id)
	if err != nil {
		return nil, fmt.Errorf("get prompt by id: %w", err)
	}
	prompt, err := pgx.CollectOneRow(rows, r.mapPrompts)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("map prompt: %w", err)
	}
	return &prompt, nil
}

func (r *Repo) GetPromptByUserQuestionAndType(ctx context.Context, prompt domain.Prompt) (*domain.Prompt, error) {
	rows, err := r.pool.GetTx(ctx).Query(ctx, getPromptByUserQuestionAndTypeQuery, prompt.UserId, prompt.Question, prompt.Type)
	if err != nil {
		return nil, fmt.Errorf("get prompt by id: %w", err)
	}
	prompt, err = pgx.CollectOneRow(rows, r.mapPrompts)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("map prompt: %w", err)
	}
	return &prompt, nil
}

func (r *Repo) UpdatePromptsPositions(ctx context.Context, prompts []domain.Prompt) ([]domain.Prompt, error) {
	//TODO implement me
	panic("implement me")
}
