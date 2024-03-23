package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/soulmate-dating/profiles.git/internal/models"
)

type Repo struct {
	pool        *pgxpool.Pool
	mapProfiles func(row pgx.CollectableRow) (models.Profile, error)
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{
		pool:        pool,
		mapProfiles: pgx.RowToStructByName[models.Profile],
	}
}

func (r *Repo) CreateProfile(ctx context.Context, p *models.Profile) error {
	var args []any
	args = append(args,
		p.UserId, p.FirstName, p.LastName, p.BirthDate,
		p.Sex, p.PreferredPartner, p.Intention, p.Height,
		p.HasChildren, p.FamilyPlans, p.Location, p.EducationLevel,
		p.DrinksAlcohol, p.SmokesCigarettes,
	)

	if _, err := r.pool.Exec(ctx, createProfileQuery, args...); err != nil {
		return fmt.Errorf("create profile: %w", err)
	}

	return nil
}

func (r *Repo) GetProfileByID(ctx context.Context, id string) (models.Profile, error) {
	rows, err := r.pool.Query(ctx, getProfileByIDQuery, id)
	if err != nil {
		return models.Profile{}, fmt.Errorf("get profile by id: %w", err)
	}
	profile, err := pgx.CollectOneRow(rows, r.mapProfiles)
	if err != nil {
		return models.Profile{}, fmt.Errorf("map profile: %w", err)
	}
	return profile, nil
}
