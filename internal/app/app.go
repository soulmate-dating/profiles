package app

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/soulmate-dating/profiles/internal/adapters/postgres"
	"github.com/soulmate-dating/profiles/internal/app/clients/media"
	"github.com/soulmate-dating/profiles/internal/config"
	"github.com/soulmate-dating/profiles/internal/domain"
	"log"
)

type App interface {
	CreateProfile(ctx context.Context, profile *domain.Profile) (*domain.Profile, error)
	GetProfile(ctx context.Context, userId uuid.UUID) (*domain.Profile, error)
	UpdateProfile(ctx context.Context, profile domain.Profile) (*domain.Profile, error)
	GetRandomProfilePreferredByUser(ctx context.Context, userId uuid.UUID) (*domain.FullProfile, error)
	GetFullProfile(ctx context.Context, userId uuid.UUID) (*domain.FullProfile, error)

	GetPrompts(ctx context.Context, userId uuid.UUID) ([]domain.Prompt, error)
	AddPrompts(ctx context.Context, prompts []domain.Prompt) ([]domain.Prompt, error)
	UpdatePrompt(ctx context.Context, prompt domain.Prompt) (*domain.Prompt, error)
	UpdatePromptsPositions(ctx context.Context, prompts []domain.Prompt) ([]domain.Prompt, error)
	GetMultipleProfiles(ctx context.Context, ids []uuid.UUID) ([]domain.Profile, error)
	AddFilePrompt(ctx context.Context, prompt domain.FilePrompt) (*domain.Prompt, error)
	UpdateFilePrompt(ctx context.Context, prompt domain.FilePrompt) (*domain.Prompt, error)
}

type Repository interface {
	CreateProfile(ctx context.Context, p *domain.Profile) error
	GetProfileByID(ctx context.Context, id uuid.UUID) (*domain.Profile, error)
	UpdateProfile(ctx context.Context, profile domain.Profile) (*domain.Profile, error)
	GetMultipleProfilesByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.Profile, error)
	GetRandomProfileBySexAndPreference(
		ctx context.Context, requesterId uuid.UUID, preference domain.Preference, sex string,
	) (*domain.Profile, error)

	GetPromptsByUser(ctx context.Context, userId uuid.UUID) ([]domain.Prompt, error)
	GetPromptByID(ctx context.Context, id uuid.UUID) (*domain.Prompt, error)
	GetPromptByUserQuestionAndType(ctx context.Context, prompt domain.Prompt) (*domain.Prompt, error)
	CreatePrompt(ctx context.Context, prompt domain.Prompt) error
	UpdatePromptContent(ctx context.Context, prompt domain.Prompt) (*domain.Prompt, error)
	UpdatePromptsPositions(ctx context.Context, prompts []domain.Prompt) ([]domain.Prompt, error)
}

type TransactionManager interface {
	RunInTx(ctx context.Context, f func(ctx context.Context) error) error
}

type Application struct {
	validate    *validator.Validate
	txManager   TransactionManager
	repository  Repository
	mediaClient media.MediaServiceClient
}

func (a *Application) UpdateFilePrompt(ctx context.Context, filePrompt domain.FilePrompt) (res *domain.Prompt, err error) {
	err = a.validate.Struct(filePrompt)
	if err != nil {
		return nil, fmt.Errorf("invalid file prompt: %w", err)
	}
	response, err := a.mediaClient.UploadFile(ctx, &media.UploadFileRequest{
		ContentType: "image/png",
		Data:        filePrompt.Content,
	})
	if err != nil {
		return nil, err
	}

	prompt := domain.Prompt{
		ID:       filePrompt.ID,
		UserId:   filePrompt.UserId,
		Question: filePrompt.Question,
		Content:  response.GetLink(),
		Position: filePrompt.Position,
		Type:     filePrompt.Type,
	}
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		res, err = a.updatePrompt(ctx, prompt)
		if err != nil {
			return fmt.Errorf("failed to update prompt: %w", err)
		}
		return nil
	})
	return res, err
}

func (a *Application) AddFilePrompt(ctx context.Context, filePrompt domain.FilePrompt) (prompt *domain.Prompt, err error) {
	err = a.validate.Struct(filePrompt)
	if err != nil {
		return nil, fmt.Errorf("invalid file prompt: %w", err)
	}
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		prompt, err = a.addFilePrompt(ctx, filePrompt)
		if err != nil {
			return fmt.Errorf("failed to add file prompt: %w", err)
		}
		return nil
	})
	return prompt, err
}

func (a *Application) addFilePrompt(ctx context.Context, filePrompt domain.FilePrompt) (*domain.Prompt, error) {
	_, err := a.repository.GetProfileByID(ctx, filePrompt.UserId)
	if err != nil {
		return nil, domain.ErrAddPromptsOnEmptyProfile
	}
	response, err := a.mediaClient.UploadFile(ctx, &media.UploadFileRequest{
		ContentType: "image/png",
		Data:        filePrompt.Content,
	})
	if err != nil {
		return nil, err
	}

	prompt := domain.Prompt{
		ID:       domain.NewUID(),
		UserId:   filePrompt.UserId,
		Question: filePrompt.Question,
		Content:  response.GetLink(),
		Position: filePrompt.Position,
		Type:     filePrompt.Type,
	}
	err = a.addPrompt(ctx, prompt)
	if err != nil {
		return nil, err
	}
	return &prompt, err
}

func (a *Application) GetFullProfile(ctx context.Context, userId uuid.UUID) (profile *domain.FullProfile, err error) {
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		profile, err = a.getFullProfile(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get full profile: %w", err)
		}
		return nil
	})
	return profile, err
}

func (a *Application) getFullProfile(ctx context.Context, userId uuid.UUID) (*domain.FullProfile, error) {
	profile, err := a.getProfile(ctx, userId)
	if err != nil {
		return nil, err
	}

	prompts, err := a.repository.GetPromptsByUser(ctx, profile.UserId)
	if err != nil {
		return nil, fmt.Errorf("get prompts: %w", err)
	}

	return &domain.FullProfile{
		Profile: *profile,
		Prompts: prompts,
	}, nil
}

func (a *Application) GetRandomProfilePreferredByUser(ctx context.Context, userId uuid.UUID) (profile *domain.FullProfile, err error) {
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		profile, err = a.getRandomProfilePreferredByUser(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get recommendation: %w", err)
		}
		return nil
	})
	return profile, err
}

func (a *Application) getRandomProfilePreferredByUser(ctx context.Context, userId uuid.UUID) (*domain.FullProfile, error) {
	profile, err := a.repository.GetProfileByID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("get profile: %w", err)
	}

	recommendedProfile, err := a.repository.GetRandomProfileBySexAndPreference(
		ctx, profile.UserId, domain.Preference(profile.PreferredPartner), profile.Sex,
	)
	if err != nil {
		return nil, fmt.Errorf("get recommedation: %w", err)
	}

	prompts, err := a.repository.GetPromptsByUser(ctx, recommendedProfile.UserId)
	if err != nil {
		return nil, fmt.Errorf("get prompt for recommended profile: %w", err)
	}

	return &domain.FullProfile{
		Profile: *recommendedProfile,
		Prompts: prompts,
	}, nil
}

func (a *Application) GetMultipleProfiles(ctx context.Context, ids []uuid.UUID) ([]domain.Profile, error) {
	p, err := a.repository.GetMultipleProfilesByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *Application) CreateProfile(ctx context.Context, profile *domain.Profile) (res *domain.Profile, err error) {
	err = a.validate.Struct(profile)
	if err != nil {
		return nil, fmt.Errorf("invalid profile: %w", err)
	}

	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		res, err = a.createProfile(ctx, profile)
		if err != nil {
			return fmt.Errorf("failed to create profile: %w", err)
		}
		return nil
	})
	return res, err
}

func (a *Application) createProfile(ctx context.Context, profile *domain.Profile) (*domain.Profile, error) {
	_, err := a.repository.GetProfileByID(ctx, profile.UserId)
	if err == nil {
		return nil, domain.ErrIDAlreadyExists
	}
	err = a.repository.CreateProfile(ctx, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (a *Application) GetProfile(ctx context.Context, userId uuid.UUID) (profile *domain.Profile, err error) {
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		profile, err = a.getProfile(ctx, userId)
		if err != nil {
			return fmt.Errorf("failed to get profile: %w", err)
		}
		return nil
	})
	return profile, err
}

func (a *Application) getProfile(ctx context.Context, userId uuid.UUID) (*domain.Profile, error) {
	p, err := a.repository.GetProfileByID(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("get profile: %w", err)
	}
	if p.MainPicPromptID != nil {
		prompt, err := a.repository.GetPromptByID(ctx, *p.MainPicPromptID)
		if err != nil {
			return nil, fmt.Errorf("get prompt for profile pic: %w", err)
		}
		p.MainPicLink = prompt.Content
	}

	return p, nil
}

func (a *Application) UpdateProfile(ctx context.Context, profile domain.Profile) (res *domain.Profile, err error) {
	err = a.validate.Struct(profile)
	if err != nil {
		return nil, fmt.Errorf("invalid profile: %w", err)
	}
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		res, err = a.updateProfile(ctx, profile)
		if err != nil {
			return fmt.Errorf("failed to update profile: %w", err)
		}
		return nil
	})
	return res, err
}

func (a *Application) updateProfile(ctx context.Context, profile domain.Profile) (*domain.Profile, error) {
	p, err := a.repository.GetProfileByID(ctx, profile.UserId)
	if err != nil {
		return nil, fmt.Errorf("get profile: %w", err)
	}

	if p.UserId.String() != profile.UserId.String() {
		return nil, domain.ErrForbidden
	}

	p, err = a.repository.UpdateProfile(ctx, profile)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (a *Application) GetPrompts(ctx context.Context, userId uuid.UUID) ([]domain.Prompt, error) {
	prompts, err := a.repository.GetPromptsByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return prompts, nil
}

func (a *Application) AddPrompts(ctx context.Context, prompts []domain.Prompt) (res []domain.Prompt, err error) {
	for _, prompt := range prompts {
		err = a.validate.Struct(prompt)
		if err != nil {
			return nil, fmt.Errorf("invalid prompt: %w", err)
		}
	}

	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		res, err = a.addPrompts(ctx, prompts)
		if err != nil {
			return fmt.Errorf("failed to add prompts: %w", err)
		}
		return nil
	})
	return res, err
}

func (a *Application) addPrompts(ctx context.Context, prompts []domain.Prompt) ([]domain.Prompt, error) {
	_, err := a.repository.GetProfileByID(ctx, prompts[0].UserId)
	if err != nil {
		return nil, domain.ErrAddPromptsOnEmptyProfile
	}
	for i := range prompts {
		prompts[i].ID = domain.NewUID()
		err := a.addPrompt(ctx, prompts[i])
		if err != nil {
			return nil, err
		}
	}

	return prompts, nil
}

func (a *Application) addPrompt(ctx context.Context, prompt domain.Prompt) error {
	_, err := a.repository.GetPromptByID(ctx, prompt.ID)
	if err == nil {
		return domain.ErrIDAlreadyExists
	}

	_, err = a.repository.GetPromptByUserQuestionAndType(ctx, prompt)
	if err == nil {
		return domain.ErrNotUnique
	}

	err = a.repository.CreatePrompt(ctx, prompt)
	if err != nil {
		return fmt.Errorf("create prompt: %w", err)
	}

	var profile *domain.Profile
	if prompt.Type == domain.Image {
		profile, err = a.repository.GetProfileByID(ctx, prompt.UserId)
		if err != nil {
			return fmt.Errorf("get profile: %w", err)
		}
		if profile.MainPicPromptID == nil {
			profile.MainPicPromptID = &prompt.ID
			_, err = a.repository.UpdateProfile(ctx, *profile)
			if err != nil {
				return fmt.Errorf("update profile: %w", err)
			}
		}
	}
	return nil
}

func (a *Application) UpdatePrompt(ctx context.Context, prompt domain.Prompt) (res *domain.Prompt, err error) {
	err = a.validate.Struct(prompt)
	if err != nil {
		return nil, fmt.Errorf("invalid prompt: %w", err)
	}
	err = a.txManager.RunInTx(ctx, func(ctx context.Context) error {
		res, err = a.updatePrompt(ctx, prompt)
		if err != nil {
			return fmt.Errorf("failed to update prompt: %w", err)
		}
		return nil
	})
	return res, err
}

func (a *Application) updatePrompt(ctx context.Context, prompt domain.Prompt) (*domain.Prompt, error) {
	p, err := a.repository.GetPromptByID(ctx, prompt.ID)
	if err != nil {
		return nil, fmt.Errorf("get prompt: %w", err)
	}

	if p.UserId.String() != prompt.UserId.String() {
		return nil, domain.ErrForbidden
	}

	p, err = a.repository.GetPromptByUserQuestionAndType(ctx, prompt)
	if err == nil && p.ID.String() != prompt.ID.String() {
		return nil, domain.ErrNotUnique
	}

	p, err = a.repository.UpdatePromptContent(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("update prompt: %w", err)
	}

	return p, nil
}

func (a *Application) UpdatePromptsPositions(ctx context.Context, prompts []domain.Prompt) ([]domain.Prompt, error) {
	prompts, err := a.repository.UpdatePromptsPositions(ctx, prompts)
	if err != nil {
		return nil, err
	}

	return prompts, nil
}

func New(ctx context.Context, cfg config.Config) App {
	conn, err := postgres.Connect(ctx, postgres.Config{
		Host:              cfg.Postgres.Host,
		Port:              cfg.Postgres.Port,
		User:              cfg.Postgres.User,
		Password:          cfg.Postgres.Password,
		DBName:            cfg.Postgres.Database,
		SSLMode:           cfg.Postgres.SSLMode,
		ConnectionTimeout: cfg.Postgres.ConnectionTimeout,
	})
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	pool := postgres.NewPool(conn)
	repo := postgres.NewRepo(pool)

	mediaClient, err := media.NewServiceClient(media.Config{
		Address:   cfg.Media.Address,
		EnableTLS: cfg.Media.EnableTLS,
	})
	if err != nil {
		log.Fatalf("could not connect to media service: %s", err.Error())
	}
	return &Application{repository: repo, mediaClient: mediaClient, txManager: pool, validate: validator.New()}
}
