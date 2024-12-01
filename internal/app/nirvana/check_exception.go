package nirvana

import (
	"context"
	"fmt"
	"github.com/samber/lo"
	"nirvana/internal/app/nirvana/usecases/model"
	"nirvana/pkg/api/nirvana"
)

func (e ExceptionService) CheckException(ctx context.Context, req *nirvana.CheckExceptionRequest) (*nirvana.CheckExceptionResponse, error) {
	fmt.Printf("[CheckException] req: %+v\n", req)

	found, err := e.checkExceptionUseCase.CheckException(ctx, req.Name, model.ExceptionAttributes{
		ClientID: lo.FromPtr(req.GetAttributes().ClientId),
		Amount:   lo.FromPtr(req.GetAttributes().Amount),
	})

	if err != nil {
		return &nirvana.CheckExceptionResponse{
			Found: false,
		}, err
	}

	return &nirvana.CheckExceptionResponse{
		Found: found,
	}, nil

}
