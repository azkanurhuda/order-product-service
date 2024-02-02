package product

import (
	"context"
	"fmt"
	"github.com/azkanurhuda/product-service/internal/entity"
	"github.com/azkanurhuda/product-service/internal/model"
	"github.com/azkanurhuda/product-service/internal/model/converter"
	"github.com/azkanurhuda/product-service/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type ProductUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (p ProductUseCase) Create(ctx context.Context, request *model.ProductRequest) (*model.ProductResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	ID := 0

	err := p.Validate.Struct(request)
	if err != nil {
		p.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, _ := p.Repository.Product.GetLatestProduct(tx)
	fmt.Println(data)
	if data != nil {
		ID += data.ID + 1
	}

	product := &entity.Product{
		ID:        ID,
		Name:      request.Name,
		Price:     request.Price,
		Stock:     request.Stock,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	if err := p.Repository.Product.Create(tx, product); err != nil {
		p.Log.Warnf("Failed create product to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}

func (p ProductUseCase) ReadByID(ctx context.Context, id int) (*model.ProductResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := p.Repository.Product.ReadByID(tx, id)
	if data == nil {
		p.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(data), nil
}

func (p ProductUseCase) ReadByAll(ctx context.Context) ([]model.ProductResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := p.Repository.Product.ReadAll(tx)
	if data == nil {
		p.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if len(data) == 0 {
		p.Log.Warn("No products found")
		tx.Rollback()
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListProductToResponse(data), nil
}

func (p ProductUseCase) Update(ctx context.Context, request *model.ProductRequest, ID int) (*model.ProductResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.Validate.Struct(request)
	if err != nil {
		p.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	product := &entity.Product{
		ID:        ID,
		Name:      request.Name,
		Price:     request.Price,
		Stock:     request.Stock,
		UpdatedAt: time.Now().Local(),
	}

	if err := p.Repository.Product.Update(tx, product); err != nil {
		p.Log.Warnf("Failed create product to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}

func (p ProductUseCase) Delete(ctx context.Context, productID int) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Repository.Product.Delete(tx, productID); err != nil {
		p.Log.Warnf("Failed create product to database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func NewProductUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &ProductUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
