package nirvana

import (
	"context"
	"github.com/samber/lo"
	"nirvana/internal/app/nirvana/usecases/model"
	"nirvana/pkg/api/nirvana"
)

func (e ExceptionService) CreateException(ctx context.Context, req *nirvana.CreateExceptionRequest) (*nirvana.CreateExceptionResponse, error) {
	err := e.createExceptionUseCase.CreateException(ctx, req.Name, model.ExceptionAttributes{
		ClientID: lo.FromPtr(req.GetAttributes().ClientId),
		Amount:   lo.FromPtr(req.GetAttributes().Amount),
	})

	if err != nil {
		return nil, err
	}

	return &nirvana.CreateExceptionResponse{}, nil
}
