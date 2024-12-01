package usecases

import (
	"context"
	"fmt"
	"nirvana/internal/app/nirvana/usecases/model"
)

type CreateExceptionRepository interface {
	CreateException(ctx context.Context, name string, attributes model.ExceptionAttributes) error
}

type CreateExceptionUseCase struct {
	repo CreateExceptionRepository
}

func NewCreateExceptionUseCase(repo CreateExceptionRepository) *CreateExceptionUseCase {
	return &CreateExceptionUseCase{
		repo: repo,
	}
}

func (uc *CreateExceptionUseCase) CreateException(ctx context.Context, name string, attributes model.ExceptionAttributes) error {
	if name == "" {
		return fmt.Errorf("exception name cannot be empty")
	}

	err := uc.repo.CreateException(ctx, name, attributes)
	if err != nil {
		return fmt.Errorf("failed to create exception: %w", err)
	}

	return nil
}
