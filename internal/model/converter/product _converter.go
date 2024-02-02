package converter

import (
	"github.com/azkanurhuda/product-service/internal/entity"
	"github.com/azkanurhuda/product-service/internal/model"
	"time"
)

func ProductToResponse(data *entity.Product) *model.ProductResponse {
	return &model.ProductResponse{
		ID:        data.ID,
		Name:      data.Name,
		Price:     data.Price,
		Stock:     data.Stock,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func ListProductToResponse(data []entity.Product) []model.ProductResponse {
	list := make([]model.ProductResponse, 0)
	for _, v := range data {
		list = append(list, model.ProductResponse{
			ID:        v.ID,
			Name:      v.Name,
			Price:     v.Price,
			Stock:     v.Stock,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return list
}
