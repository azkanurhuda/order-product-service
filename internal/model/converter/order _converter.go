package converter

import (
	"github.com/azkanurhuda/order-product-service/internal/entity"
	"github.com/azkanurhuda/order-product-service/internal/model"
	"github.com/azkanurhuda/order-product-service/sdk/http/costumer_sdk"
	"github.com/azkanurhuda/order-product-service/sdk/http/product_sdk"
	"time"
)

func OrderToResponse(data *entity.Order) *model.OrderResponse {
	return &model.OrderResponse{
		ID:         data.ID,
		CustomerID: data.CustomerID,
		ProductID:  data.ProductID,
		Quantity:   data.Quantity,
		Total:      data.Total,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}
}

func OrderToFullResponse(data *entity.Order, customerResponse *costumer_sdk.CostumerResponse, productResponse *product_sdk.ProductResponse) *model.OrderFullResponse {
	response := &model.OrderFullResponse{
		ID: data.ID,
		Customer: model.CostumerResponse{
			ID:        customerResponse.ID,
			Name:      customerResponse.Name,
			Email:     customerResponse.Email,
			CreatedAt: customerResponse.CreatedAt,
			UpdatedAt: customerResponse.UpdatedAt,
		},
		Product: model.ProductResponse{
			ID:        productResponse.ID,
			Name:      productResponse.Name,
			Price:     productResponse.Price,
			Stock:     productResponse.Stock,
			CreatedAt: productResponse.CreatedAt,
			UpdatedAt: productResponse.UpdatedAt,
		},
		Quantity:  data.Quantity,
		Total:     data.Total,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return response
}

func ListOrderToResponse(data []entity.Order) []model.OrderResponse {
	list := make([]model.OrderResponse, 0)
	for _, v := range data {
		list = append(list, model.OrderResponse{
			ID:         v.ID,
			CustomerID: v.CustomerID,
			ProductID:  v.ProductID,
			Quantity:   v.Quantity,
			Total:      v.Total,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
		})
	}
	return list
}
