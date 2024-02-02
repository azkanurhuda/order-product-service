package product

import (
	"context"
	"github.com/azkanurhuda/product-service/internal/model"
)

type UseCase interface {
	Create(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error)
	ReadByID(ctx context.Context, id int) (*model.ProductResponse, error)
	ReadByAll(ctx context.Context) ([]model.ProductResponse, error)
	Update(ctx context.Context, request *model.ProductRequest, ID int) (*model.ProductResponse, error)
	Delete(ctx context.Context, productID int) error
}
