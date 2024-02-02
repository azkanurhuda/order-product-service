package costumer

import (
	"context"
	"github.com/azkanurhuda/costumer-service/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.CostumerRequest) (*model.CostumerResponse, error)
	ReadByID(ctx context.Context, id int) (*model.CostumerResponse, error)
	ReadByAll(ctx context.Context) ([]model.CostumerResponse, error)
	Update(ctx context.Context, request *model.CostumerRequest, ID int) (*model.CostumerResponse, error)
	Delete(ctx context.Context, ID int) error
}
