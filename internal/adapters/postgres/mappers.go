package postgres

import (
	"github.com/google/uuid"
	"github.com/soulmate-dating/profiles/internal/domain"
)

type PromptBatch struct {
	IDs       []uuid.UUID
	Positions []int32
}

func NewPromptBatch(prompts []domain.Prompt) PromptBatch {
	ids := make([]uuid.UUID, len(prompts))
	positions := make([]int32, len(prompts))
	for i, p := range prompts {
		ids[i] = p.ID
		positions[i] = p.Position
	}
	return PromptBatch{
		IDs:       ids,
		Positions: positions,
	}
}
