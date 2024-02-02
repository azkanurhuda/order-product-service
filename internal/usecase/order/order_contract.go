package order

import (
	"context"
	"github.com/azkanurhuda/order-product-service/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.OrderRequest) (*model.OrderResponse, error)
	ReadByID(ctx context.Context, id int) (*model.OrderFullResponse, error)
	ReadByAll(ctx context.Context) ([]model.OrderResponse, error)
	Update(ctx context.Context, request *model.OrderRequest, ID int) (*model.OrderResponse, error)
	Delete(ctx context.Context, ID int) error
}
