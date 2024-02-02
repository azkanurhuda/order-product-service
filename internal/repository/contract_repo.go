package repository

import (
	"github.com/azkanurhuda/costumer-service/internal/entity"
	"gorm.io/gorm"
)

type Costumer interface {
	Create(db *gorm.DB, product *entity.Costumer) error
	ReadByID(db *gorm.DB, id int) (*entity.Costumer, error)
	ReadAll(db *gorm.DB) ([]entity.Costumer, error)
	Update(db *gorm.DB, product *entity.Costumer) error
	Delete(db *gorm.DB, productID int) error
	GetLatestCostumer(db *gorm.DB) (*entity.Costumer, error)
}
