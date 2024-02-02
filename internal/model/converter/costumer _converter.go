package converter

import (
	"github.com/azkanurhuda/costumer-service/internal/entity"
	"github.com/azkanurhuda/costumer-service/internal/model"
	"time"
)

func CostumerToResponse(data *entity.Costumer) *model.CostumerResponse {
	return &model.CostumerResponse{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func ListCostumerToResponse(data []entity.Costumer) []model.CostumerResponse {
	list := make([]model.CostumerResponse, 0)
	for _, v := range data {
		list = append(list, model.CostumerResponse{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return list
}
