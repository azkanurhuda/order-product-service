package costumer

import (
	"context"
	"fmt"
	"github.com/azkanurhuda/costumer-service/internal/entity"
	"github.com/azkanurhuda/costumer-service/internal/model"
	"github.com/azkanurhuda/costumer-service/internal/model/converter"
	"github.com/azkanurhuda/costumer-service/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type CostumerUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (p CostumerUseCase) Create(ctx context.Context, request *model.CostumerRequest) (*model.CostumerResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	ID := 0

	err := p.Validate.Struct(request)
	if err != nil {
		p.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, _ := p.Repository.Costumer.GetLatestCostumer(tx)
	fmt.Println(data)
	if data != nil {
		ID += data.ID + 1
	}

	product := &entity.Costumer{
		ID:        ID,
		Name:      request.Name,
		Email:     request.Email,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	if err := p.Repository.Costumer.Create(tx, product); err != nil {
		p.Log.Warnf("Failed create costumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CostumerToResponse(product), nil
}

func (p CostumerUseCase) ReadByID(ctx context.Context, id int) (*model.CostumerResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := p.Repository.Costumer.ReadByID(tx, id)
	if data == nil {
		p.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CostumerToResponse(data), nil
}

func (p CostumerUseCase) ReadByAll(ctx context.Context) ([]model.CostumerResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data, err := p.Repository.Costumer.ReadAll(tx)
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

	return converter.ListCostumerToResponse(data), nil
}

func (p CostumerUseCase) Update(ctx context.Context, request *model.CostumerRequest, ID int) (*model.CostumerResponse, error) {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := p.Validate.Struct(request)
	if err != nil {
		p.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	product := &entity.Costumer{
		ID:        ID,
		Name:      request.Name,
		Email:     request.Email,
		UpdatedAt: time.Now().Local(),
	}

	if err := p.Repository.Costumer.Update(tx, product); err != nil {
		p.Log.Warnf("Failed create costumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CostumerToResponse(product), nil
}

func (p CostumerUseCase) Delete(ctx context.Context, ID int) error {
	tx := p.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := p.Repository.Costumer.Delete(tx, ID); err != nil {
		p.Log.Warnf("Failed create costumer to database : %+v", err)
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.Warnf("Failed commit transaction : %+v", err)
		return fiber.ErrInternalServerError
	}

	return nil
}

func NewCostumerUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &CostumerUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
