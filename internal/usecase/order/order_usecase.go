package order

import (
	"context"
	"fmt"
	"github.com/azkanurhuda/order-product-service/internal/entity"
	"github.com/azkanurhuda/order-product-service/internal/model"
	"github.com/azkanurhuda/order-product-service/internal/model/converter"
	"github.com/azkanurhuda/order-product-service/internal/repository"
	"github.com/azkanurhuda/order-product-service/sdk/http/costumer_sdk"
	"github.com/azkanurhuda/order-product-service/sdk/http/product_sdk"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type OrderUseCase struct {
	DB          *gorm.DB
	Log         *logrus.Logger
	Validate    *validator.Validate
	Repository  *repository.Repository
	CostumerSDK costumer_sdk.SDK
	ProductSDK  product_sdk.SDK
}

func (p OrderUseCase) Create(ctx context.Context, request *model.OrderRequest) (*model.OrderResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	ID := 0

	err := p.Validate.Struct(request)
	if err != nil {
		p.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, _ := p.Repository.Order.GetLatestOrder(tx)
	fmt.Println(data)
	if data != nil {
		ID += data.ID + 1
	}

	_, err = p.CostumerSDK.GetCustomerByID(ctx, request.CustomerID)
	fmt.Println("err 1", err)
	if err != nil {
		p.Log.Warnf("Costumer not exists : %+v", err)
		return nil, fiber.ErrNotFound

	}

	_, err = p.ProductSDK.GetProductByID(ctx, request.ProductID)
	fmt.Println("err 2", err)
	if err != nil {
		p.Log.Warnf("Product not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	order := &entity.Order{
		ID:         ID,
		CustomerID: request.CustomerID,
		ProductID:  request.ProductID,
		Quantity:   request.Quantity,
		Total:      request.Total,
		CreatedAt:  time.Now().Local(),
		UpdatedAt:  time.Now().Local(),
	}

	if err := p.Repository.Order.Create(tx, order); err != nil {
		p.Log.Warnf("Failed create order to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}

func (p OrderUseCase) ReadByID(ctx context.Context, id int) (*model.OrderFullResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := p.Repository.Order.ReadByID(tx, id)
	if data == nil {
		p.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	responseConsumer, err := p.CostumerSDK.GetCustomerByID(ctx, data.CustomerID)
	if err != nil {
		p.Log.Warnf("Costumer not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	responseProduct, err := p.ProductSDK.GetProductByID(ctx, data.ProductID)
	if err != nil {
		p.Log.Warnf("Costumer not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToFullResponse(data, responseConsumer, responseProduct), nil
}

func (p OrderUseCase) ReadByAll(ctx context.Context) ([]model.OrderResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := p.Repository.Order.ReadAll(tx)
	if data == nil {
		p.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if len(data) == 0 {
		p.Log.Warn("No orders found")
		tx.Rollback()
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListOrderToResponse(data), nil
}

func (p OrderUseCase) Update(ctx context.Context, request *model.OrderRequest, ID int) (*model.OrderResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.Validate.Struct(request)
	if err != nil {
		p.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	order := &entity.Order{
		ID:         ID,
		CustomerID: request.CustomerID,
		ProductID:  request.ProductID,
		Quantity:   request.Quantity,
		Total:      request.Total,
		UpdatedAt:  time.Now().Local(),
	}

	if err := p.Repository.Order.Update(tx, order); err != nil {
		p.Log.Warnf("Failed create order to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.OrderToResponse(order), nil
}

func (p OrderUseCase) Delete(ctx context.Context, ID int) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Repository.Order.Delete(tx, ID); err != nil {
		p.Log.Warnf("Failed create order to database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func NewOrderUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository, costumerSDK costumer_sdk.SDK, productSDK product_sdk.SDK) UseCase {
	return &OrderUseCase{
		DB:          db,
		Log:         logger,
		Validate:    validate,
		Repository:  repo,
		CostumerSDK: costumerSDK,
		ProductSDK:  productSDK,
	}
}
