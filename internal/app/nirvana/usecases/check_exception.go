package usecases

import (
	"context"
	"fmt"
	"nirvana/internal/app/nirvana/repository"
	"nirvana/internal/app/nirvana/usecases/model"
)

type CheckExceptionRepository interface {
	CheckException(ctx context.Context, name string, attributes model.ExceptionAttributes) (bool, error)
}

type CheckExceptionUseCase struct {
	repo *repository.Repository
}

func NewCheckExceptionUseCase(repo *repository.Repository) *CheckExceptionUseCase {
	return &CheckExceptionUseCase{
		repo: repo,
	}
}

func (uc *CheckExceptionUseCase) CheckException(ctx context.Context, name string, attributes model.ExceptionAttributes) (bool, error) {
	if name == "" {
		return false, fmt.Errorf("exception name cannot be empty")
	}

	found, err := uc.repo.CheckException(ctx, name, attributes)
	if err != nil {
		return false, fmt.Errorf("failed to check exception: %w", err)
	}

	return found, nil
}
