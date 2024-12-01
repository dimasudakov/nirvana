package nirvana

import (
	"context"
	"nirvana/internal/app/nirvana/usecases/model"
	"nirvana/pkg/api/nirvana"
)

type CreateExceptionUseCase interface {
	CreateException(ctx context.Context, name string, attributes model.ExceptionAttributes) error
}

type CheckExceptionUseCase interface {
	CheckException(ctx context.Context, name string, attributes model.ExceptionAttributes) (bool, error)
}

type ExceptionService struct {
	createExceptionUseCase CreateExceptionUseCase
	checkExceptionUseCase  CheckExceptionUseCase

	nirvana.UnimplementedNirvanaServer
}

func NewService(
	createExceptionUseCase CreateExceptionUseCase,
	checkExceptionUseCase CheckExceptionUseCase,
) *ExceptionService {
	return &ExceptionService{
		createExceptionUseCase: createExceptionUseCase,
		checkExceptionUseCase:  checkExceptionUseCase,
	}
}
